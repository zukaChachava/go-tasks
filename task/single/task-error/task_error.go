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

func (task *TaskErr[TErr]) Run() *TaskResult[TErr] {
	channel := make(chan TErr, 1)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(channel chan TErr, wg *sync.WaitGroup) {
		defer wg.Done()
		err := task.task()
		channel <- err
	}(channel, &wg)

	return &TaskResult[TErr]{
		wait:    &wg,
		channel: channel,
	}
}

type TaskResult[TErr error] struct {
	wait    *sync.WaitGroup
	channel chan TErr
}

func (task *TaskResult[TErr]) Wait() TErr {
	task.wait.Wait()
	return <-task.channel
}
