package multiple

import (
	tasks "github.com/zukaChachava/task/task/multiple/tasks-result-error"
	"testing"
	"time"
)

func Test_RunWithResultError(t *testing.T) {
	tasksContainer := tasks.NewTasks[int, error]()

	for i := 0; i < 10; i++ {
		tasksContainer.Add(func() (int, error) {
			time.Sleep(1 * time.Second)
			return i, nil
		})
	}

	currentTime := time.Now()
	results := tasksContainer.Run().Wait()
	duration := time.Since(currentTime)

	num := 0
	for _, result := range results {
		if nil != result.Err {
			t.Fatal("Invalid result")
		}
		num++
	}

	if duration > time.Second*2 {
		t.Fatal("Not running concurrently")
	}
}
