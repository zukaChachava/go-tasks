package multiple

import (
	tasks "github.com/zukaChachava/task/task/multiple/tasks-result"
	"testing"
	"time"
)

func Test_RunWithResult_Returns10(t *testing.T) {
	tasksContainer := tasks.NewTasks[int]()

	for i := 0; i < 10; i++ {
		tasksContainer.Add(func() int {
			time.Sleep(1 * time.Second)
			return i
		})
	}

	currentTime := time.Now()
	results := tasksContainer.Run().Wait()
	duration := time.Since(currentTime)

	num := 0
	for result := range results {
		if num != result {
			t.Fatal("Invalid result")
		}
		num++
	}

	if duration > time.Second*2 {
		t.Fatal("Not running concurrently")
	}
}
