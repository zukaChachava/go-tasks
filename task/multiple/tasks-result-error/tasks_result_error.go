package tasks_result_error

import "sync"

type TasksResultError[T any, TErr error] struct {
	tasks []func() (T, TErr)
}

func NewTasksResultError[T any, TErr error]() *TasksResultError[T, TErr] {
	return &TasksResultError[T, TErr]{
		tasks: make([]func() (T, TErr), 0, 4),
	}
}

func (tasks *TasksResultError[T, TErr]) Add(task func() (T, TErr)) *TasksResultError[T, TErr] {
	tasks.tasks = append(tasks.tasks, task)
	return tasks
}

func (tasks *TasksResultError[T, TErr]) Run() *TasksResult[T, TErr] {
	wg := sync.WaitGroup{}
	channel := make(chan tupleWrapper[T, TErr], len(tasks.tasks))

	for index, task := range tasks.tasks {
		wg.Add(1)
		go func(wg *sync.WaitGroup, channel chan tupleWrapper[T, TErr], task func() (T, TErr)) {
			defer wg.Done()
			result, err := task()
			channel <- tupleWrapper[T, TErr]{index, &result, err}
		}(&wg, channel, task)
	}

	return &TasksResult[T, TErr]{size: len(tasks.tasks), wg: &wg, channel: channel}
}

type TasksResult[T any, TErr error] struct {
	size    int
	wg      *sync.WaitGroup
	channel chan tupleWrapper[T, TErr]
}

func (tasks *TasksResult[T, TErr]) Wait() []ResultTuple[T, TErr] {
	tasks.wg.Wait()
	data := make([]ResultTuple[T, TErr], tasks.size)

	for result := range tasks.channel {
		data[result.index] = ResultTuple[T, TErr]{Value: result.value, Err: result.err}
	}

	return data
}

type tupleWrapper[T any, TErr error] struct {
	index int
	value *T
	err   TErr
}

type ResultTuple[T any, TErr error] struct {
	Value *T
	Err   TErr
}
