package task_no_error

import "sync"

type Task struct {
	task func()
}

func NewTask(task func()) *Task {
	return &Task{
		task: task,
	}
}

func (task *Task) Run() *TaskResult {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		task.task()
	}(&wg)

	return &TaskResult{
		wait: &wg,
	}
}

type TaskResult struct {
	wait *sync.WaitGroup
}

func (task *TaskResult) Wait() {
	task.wait.Wait()
}
