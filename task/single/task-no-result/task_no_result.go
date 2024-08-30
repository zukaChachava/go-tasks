package task_no_result

import "sync"

type Task struct {
	task func()
}

func NewTask(task func()) *Task {
	return &Task{
		task: task,
	}
}

func (task *Task) Run() *Result {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup, task func()) {
		defer wg.Done()
		task()
	}(&wg, task.task)

	return &Result{
		wait: &wg,
	}
}

type Result struct {
	wait *sync.WaitGroup
}

func (result *Result) Wait() {
	result.wait.Wait()
}
