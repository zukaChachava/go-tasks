package task_error

import (
	"sync"
)

type TaskErr[TErr error] struct {
	task func() TErr
}

func NewTask[TErr error](task func() TErr) *TaskErr[TErr] {
	return &TaskErr[TErr]{
		task: task,
	}
}

func (task *TaskErr[TErr]) Run() *Result[TErr] {
	channel := make(chan TErr, 1)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup, channel chan TErr, task func() TErr) {
		defer wg.Done()
		err := task()
		channel <- err
	}(&wg, channel, task.task)

	return &Result[TErr]{
		wait:    &wg,
		channel: channel,
	}
}

type Result[TErr error] struct {
	wait    *sync.WaitGroup
	channel chan TErr
}

func (result *Result[TErr]) Wait() TErr {
	result.wait.Wait()
	return <-result.channel
}
