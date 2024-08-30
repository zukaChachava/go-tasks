package tasks_no_error

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

func (tasks *Tasks) Run() *TasksResult {
	wg := sync.WaitGroup{}

	for _, task := range tasks.tasks {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			task()
		}(&wg)
	}

	clear(tasks.tasks)
	return &TasksResult{
		wg: &wg,
	}
}

type TasksResult struct {
	wg *sync.WaitGroup
}

func (tasksResult *TasksResult) Wait() {
	tasksResult.wg.Wait()
}
