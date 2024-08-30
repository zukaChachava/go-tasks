package task_result

import "sync"

type TaskResult[T any] struct {
	task func() T
}

func NewTask[T any](task func() T) *TaskResult[T] {
	return &TaskResult[T]{task: task}
}

func (task *TaskResult[T]) Run() *Result[T] {
	wg := sync.WaitGroup{}
	channel := make(chan T, 1)

	go func(wg *sync.WaitGroup, channel chan T, task func() T) {
		defer wg.Done()
		result := task()
		channel <- result
	}(&wg, channel, task.task)

	return &Result[T]{
		wait:    &wg,
		channel: channel,
	}
}

type Result[T any] struct {
	wait    *sync.WaitGroup
	channel chan T
}

func (result *Result[T]) Wait() *T {
	result.wait.Wait()
	data := <-result.channel
	return &data
}
