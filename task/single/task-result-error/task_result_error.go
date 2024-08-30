package task_result_error

import "sync"

type TaskResultErr[T any, TErr error] struct {
	task func() (T, TErr)
}

func NewTask[T any, TErr error](task func() (T, TErr)) *TaskResultErr[T, TErr] {
	return &TaskResultErr[T, TErr]{
		task: task,
	}
}

func (task *TaskResultErr[T, TErr]) Run() *TaskResult[T, TErr] {
	channel := make(chan tupleWrapper[T, TErr], 1)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(channel chan tupleWrapper[T, TErr], wg *sync.WaitGroup) {
		defer wg.Done()
		result, err := task.task()
		channel <- tupleWrapper[T, TErr]{result: &result, err: err}
	}(channel, &wg)

	return &TaskResult[T, TErr]{
		wait:    &wg,
		channel: channel,
	}
}

type TaskResult[T any, TErr error] struct {
	wait    *sync.WaitGroup
	channel chan tupleWrapper[T, TErr]
}

func (task *TaskResult[T, TErr]) Wait() (*T, TErr) {
	task.wait.Wait()
	data := <-task.channel
	return data.result, data.err
}

type tupleWrapper[T any, TErr error] struct {
	result *T
	err    TErr
}
