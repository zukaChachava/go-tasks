package multiple

import (
	tasks "github.com/zukaChachava/task/task/multiple/tasks-no-result"
	"testing"
	"time"
)

func Test_RunWithNoResult(t *testing.T) {
	tasksContainer := tasks.NewTasks()

	for i := 0; i < 10; i++ {
		tasksContainer.Add(func() {
			time.Sleep(1 * time.Second)
		})
	}

	currentTime := time.Now()
	tasksContainer.Run().Wait()
	duration := time.Since(currentTime)

	if duration > time.Second*2 {
		t.Fatal("Not running concurrently")
	}
}
