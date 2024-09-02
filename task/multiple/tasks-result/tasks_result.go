package tasks_result

import (
	"sync"
)

type TasksResult[T any] struct {
	tasks []func() T
}

func NewTasks[T any]() *TasksResult[T] {
	return &TasksResult[T]{
		tasks: make([]func() T, 0, 4),
	}
}

func (tasks *TasksResult[T]) Add(task func() T) *TasksResult[T] {
	tasks.tasks = append(tasks.tasks, task)
	return tasks
}

func (tasks *TasksResult[T]) Run() *Result[T] {
	wg := sync.WaitGroup{}
	channel := make(chan tupleWrapper[T], len(tasks.tasks))

	for index, task := range tasks.tasks {
		wg.Add(1)
		go func(wg *sync.WaitGroup, channel chan tupleWrapper[T], task func() T) {
			defer wg.Done()
			result := task()
			channel <- tupleWrapper[T]{index: index, value: &result}
		}(&wg, channel, task)
	}

	return &Result[T]{
		wg:      &wg,
		channel: channel,
		size:    len(tasks.tasks),
	}
}

type Result[T any] struct {
	wg      *sync.WaitGroup
	channel chan tupleWrapper[T]
	size    int
}

func (result *Result[T]) Wait() []*T {
	result.wg.Wait()
	data := make([]*T, result.size)

	for i := 0; i < result.size; i++ {
		value := <-result.channel
		data[value.index] = value.value
	}

	return data
}

type tupleWrapper[T any] struct {
	index int
	value *T
}
