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

func (task *TaskResultErr[T, TErr]) Run() *Result[T, TErr] {
	channel := make(chan tupleWrapper[T, TErr], 1)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup, channel chan tupleWrapper[T, TErr], task func() (T, TErr)) {
		defer wg.Done()
		result, err := task()
		channel <- tupleWrapper[T, TErr]{result: &result, err: err}
	}(&wg, channel, task.task)

	return &Result[T, TErr]{
		wait:    &wg,
		channel: channel,
	}
}

type Result[T any, TErr error] struct {
	wait    *sync.WaitGroup
	channel chan tupleWrapper[T, TErr]
}

func (result *Result[T, TErr]) Wait() (*T, TErr) {
	result.wait.Wait()
	data := <-result.channel
	return data.result, data.err
}

type tupleWrapper[T any, TErr error] struct {
	result *T
	err    TErr
}
