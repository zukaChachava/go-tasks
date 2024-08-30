package tasks_error

import "sync"

type TasksError[TErr error] struct {
	tasks []func() TErr
}

func NewTasks[TErr error]() *TasksError[TErr] {
	return &TasksError[TErr]{
		tasks: make([]func() TErr, 0, 4),
	}
}

func (tasks *TasksError[TErr]) Add(task func() TErr) {
	tasks.tasks = append(tasks.tasks, task)
}

func (tasks *TasksError[TErr]) Run() *Result[TErr] {
	wg := sync.WaitGroup{}
	channel := make(chan tupleWrapper[TErr], len(tasks.tasks))

	for index, task := range tasks.tasks {
		wg.Add(1)
		go func(wg *sync.WaitGroup, channel chan tupleWrapper[TErr], task func() TErr) {
			defer wg.Done()
			err := task()
			channel <- tupleWrapper[TErr]{index: index, err: err}
		}(&wg, channel, task)
	}

	return &Result[TErr]{
		wg:      &wg,
		channel: channel,
		size:    len(tasks.tasks),
	}
}

type Result[TErr error] struct {
	wg      *sync.WaitGroup
	channel chan tupleWrapper[TErr]
	size    int
}

func (result *Result[TErr]) Wait() []TErr {
	result.wg.Wait()
	results := make([]TErr, result.size)

	for data := range result.channel {
		results[data.index] = data.err
	}

	return results
}

type tupleWrapper[Err error] struct {
	index int
	err   Err
}
