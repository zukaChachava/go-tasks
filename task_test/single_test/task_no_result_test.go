package single_test

import (
	task "github.com/zukaChachava/task/task/single/task-no-result"
	"testing"
	"time"
)

func Test_RunWithNoResult(t *testing.T) {
	from := time.Now()

	taskUnit := task.NewTask(func() {
		time.Sleep(1 * time.Second)
	}).Run()

	duration1 := time.Since(from)

	if duration1 > time.Second {
		t.Fatal("duration1 should be less than 1 second")
	}

	taskUnit.Wait()

	duration2 := time.Since(from)

	if duration2 < time.Second {
		t.Fatal("duration1 should be greater than 1 second")
	}
}
