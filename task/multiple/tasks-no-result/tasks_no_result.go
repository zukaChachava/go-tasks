package tasks_no_result

import "sync"

type Tasks struct {
	tasks []func()
}

func NewTasks() *Tasks {
	return &Tasks{
		tasks: make([]func(), 0, 4),
	}
}

func (tasks *Tasks) Add(task func()) {
	tasks.tasks = append(tasks.tasks, task)
}

func (tasks *Tasks) Run() *Result {
	wg := sync.WaitGroup{}

	for _, task := range tasks.tasks {
		wg.Add(1)
		go func(wg *sync.WaitGroup, task func()) {
			defer wg.Done()
			task()
		}(&wg, task)
	}

	clear(tasks.tasks)
	return &Result{
		wg: &wg,
	}
}

type Result struct {
	wg *sync.WaitGroup
}

func (result *Result) Wait() {
	result.wg.Wait()
}
