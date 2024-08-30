package tasks_error

import "sync"

type TasksError[TErr error] struct {
	tasks []func() TErr
}

func NewTaskError[TErr error]() *TasksError[TErr] {
	return &TasksError[TErr]{
		tasks: make([]func() TErr, 0, 4),
	}
}

func (tasks *TasksError[TErr]) Add(task func() TErr) {
	tasks.tasks = append(tasks.tasks, task)
}

func (tasks *TasksError[TErr]) Run() *TasksResult[TErr] {
	wg := sync.WaitGroup{}
	channel := make(chan tupleWrapper[TErr], len(tasks.tasks))

	for index, task := range tasks.tasks {
		wg.Add(1)

		go func(channel chan tupleWrapper[TErr], wg *sync.WaitGroup) {
			defer wg.Done()
			err := task()
			channel <- tupleWrapper[TErr]{index: index, err: err}
		}(channel, &wg)
	}

	return &TasksResult[TErr]{
		wg:      &wg,
		channel: channel,
		size:    len(tasks.tasks),
	}
}

type TasksResult[TErr error] struct {
	wg      *sync.WaitGroup
	channel chan tupleWrapper[TErr]
	size    int
}

func (tasksResult *TasksResult[TErr]) Wait() []TErr {
	tasksResult.wg.Wait()
	results := make([]TErr, tasksResult.size)

	for data := range tasksResult.channel {
		results[data.index] = data.err
	}

	return results
}

type tupleWrapper[Err error] struct {
	index int
	err   Err
}
