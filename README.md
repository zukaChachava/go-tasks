# GO Tasks

This package helps write goroutines at ease. No need
to write external channels or waitgroups. Just focus
on the functionality. All the complexity is abstracted
away from devs.

The package supports executing single task as well
as executing multiple ones at the same time.

Both versions supports 4 types of execution:

* a task (tasks) without result
* a task (tasks) with a result
* a task (tasks) with a possible error
* a task (tasks) with a result and possible error
---

### Examples

- A task without result
```go
package main

import (
	task "github.com/zukaChachava/task/task/single/task-no-result"
    "time"
)

func main() {
	task.NewTask(func() {
		time.Sleep(1 * time.Second)
	}).Run().Wait()
}

```

- Tasks without result
```go
package main

import (
	tasks "github.com/zukaChachava/task/task/multiple/tasks-no-result"
    "time"
)

func main() {
	tasksContainer := tasks.NewTasks()

	for i := 0; i < 10; i++ {
		tasksContainer.Add(func() {
			time.Sleep(1 * time.Second)
		})
	}
	
	tasksContainer.Run().Wait()
```

- A task with result 
```go
package main

import (
	task "github.com/zukaChachava/task/task/single/task-result"
	"testing"
	"time"
)

func main() {
	result := task.NewTask[int](func() int {
		time.Sleep(1 * time.Second)
		return 10
	}).Run().Wait()
}
```

- Tasks with result
```go
package main

import (
	tasks "github.com/zukaChachava/task/task/multiple/tasks-result"
	"fmt"
	"time"
)

func main() {
	tasksContainer := tasks.NewTasks[int]()

	for i := 0; i < 10; i++ {
		tasksContainer.Add(func() int {
			time.Sleep(1 * time.Second)
			return i
		})
	}

	results := tasksContainer.Run().Wait()

	for result := range results {
		fmt.Println(*result)
	}
}
```

- A task with error
```go
package main

import (
	"errors"
	task "github.com/zukaChachava/task/task/single/task-error"
    "time"
)

func main() {
	err := task.NewTask[error](func() error {
		time.Sleep(1 * time.Second)
		return errors.New("test error")
	}).Run().Wait()

	if err != nil {
		panic("რაცხა არაა კაი ამბავი")
	}
}
```

- Tasks with error
```go
package main

import (
	tasks "github.com/zukaChachava/task/task/multiple/tasks-error"
    "time"
)

func main() {
	taskContainer := tasks.NewTasks[error]()

	for i := 0; i < 10; i++ {
		taskContainer.Add(func() error {
			time.Sleep(1 * time.Second)
			return nil
		})
	}
	results := taskContainer.Run().Wait()
}
```

- A task with result and error
```go
package main

import (
	task "github.com/zukaChachava/task/task/single/task-result"
	"time"
)

func main() {
	result := task.NewTask[int](func() int {
		time.Sleep(1 * time.Second)
		return 10
	}).Run().Wait()
}
```

- Tasks with result and error

```go
package main

import (
	"fmt"
	tasks "github.com/zukaChachava/task/task/multiple/tasks-result-error"
	"time"
)

func main() {
	tasksContainer := tasks.NewTasks[int, error]()

	for i := 0; i < 10; i++ {
		tasksContainer.Add(func() (int, error) {
			time.Sleep(1 * time.Second)
			return i, nil
		})
	}

	results := tasksContainer.Run().Wait()

	for _, result := range results {
		fmt.Println(*result)
	}
}

```
