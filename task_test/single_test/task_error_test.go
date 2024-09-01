package single

import (
	"errors"
	task "github.com/zukaChachava/task/task/single/task-error"
	"testing"
	"time"
)

func Test_RunWithError_ReturnsError(t *testing.T) {
	result := task.NewTask[error](func() error {
		time.Sleep(1 * time.Second)
		return errors.New("test error")
	}).Run().Wait()

	if result == nil {
		t.Fatal("Invalid result")
	}
}
