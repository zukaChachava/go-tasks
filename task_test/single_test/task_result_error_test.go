package single

import (
	task "github.com/zukaChachava/task/task/single/task-result-error"
	"testing"
	"time"
)

func TestRunWithResultError_Returns10(t *testing.T) {
	result, err := task.NewTask[int, error](func() (int, error) {
		time.Sleep(1 * time.Second)
		return 10, nil
	}).Run().Wait()

	if err != nil {
		t.Fatal(err)
	}

	if *result != 10 {
		t.Fatal("Invalid result")
	}
}
