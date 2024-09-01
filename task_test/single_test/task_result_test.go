package single_test

import (
	task "github.com/zukaChachava/task/task/single/task-result"
	"testing"
	"time"
)

func Test_RunWithResult_Returns10(t *testing.T) {
	result := task.NewTask[int](func() int {
		time.Sleep(1 * time.Second)
		return 10
	}).Run().Wait()

	if *result != 10 {
		t.Fatal("Invalid result")
	}
}
