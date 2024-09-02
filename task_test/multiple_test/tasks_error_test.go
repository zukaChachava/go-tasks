package multiple_test

import (
	tasks "github.com/zukaChachava/task/task/multiple/tasks-error"
	"testing"
	"time"
)

func Test_RunWithErrors_ReturnsNil(t *testing.T) {
	taskContainer := tasks.NewTasks[error]()

	for i := 0; i < 10; i++ {
		taskContainer.Add(func() error {
			time.Sleep(1 * time.Second)
			return nil
		})
	}

	currentTime := time.Now()
	results := taskContainer.Run().Wait()
	duration := time.Since(currentTime)

	if len(results) != 10 {
		t.Fatal("Invalid return quantity")
	}

	for _, result := range results {
		if result != nil {
			t.Fatal("Invalid result")
		}
	}

	if duration > time.Second*2 {
		t.Fatal("Not running concurrently")
	}
}
