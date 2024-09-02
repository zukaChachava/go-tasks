# GO Tasks

The GO Tasks package simplifies working with goroutines by 
abstracting away the complexity of external channels and wait groups.
Focus on the functionality of your tasks, while the package handles the 
concurrency for you. 
It supports both single and multiple tasks with four types of execution:

1. Tasks without results
2. Tasks with results
3. Tasks with possible errors
4. Tasks with results and possible errors
---

## Installation
To add the GO Tasks package to your Go module, run:

```sh
go get github.com/zukaChachava/task
````

## Usage

### Task Without Result

For tasks that do not return a result:

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

### Multiple Tasks Without Result

For running multiple tasks that do not return results:

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
}
```

### Task With Result

For a task that returns a result:

```go
package main

import (
	task "github.com/zukaChachava/task/task/single/task-result"
	"time"
	"fmt"
)

func main() {
	result := task.NewTask[int](func() int {
		time.Sleep(1 * time.Second)
		return 10
	}).Run().Wait()

	// Use the result as needed
	fmt.Println(result)
}
```

### Multiple Tasks With Results

For multiple tasks that return results:

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

	for _, result := range results {
		fmt.Println(result)
	}
}
```
### Task With Error

For a task that may return an error:

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
		panic("An error occurred: " + err.Error())
	}
}
```
### Multiple Tasks With Errors

For multiple tasks that may return errors:

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

	// Handle errors if needed
}
```

### Task With Result and Error

For a task that returns a result and may also return an error:

```go
package main

import (
	task "github.com/zukaChachava/task/task/single/task-result-error"
	"time"
	"fmt"
)

func main() {
	result, err := task.NewTask[int, error](func() (int, error) {
		time.Sleep(1 * time.Second)
		return 10, nil
	}).Run().Wait()

	if err != nil {
		panic("An error occurred: " + err.Error())
	}

	// Use the result as needed
	fmt.Println(result)
}
```

### Multiple Tasks With Results and Errors

For multiple tasks that return results and may also return errors:

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
		if result.Error != nil {
			fmt.Println("Error:", result.Error)
		} else {
			fmt.Println(result.Value)
		}
	}
}
```

## Conclusion
The GO Tasks package provides a streamlined approach to managing goroutines with various execution scenarios, making concurrency in your Go applications easier and more manageable.

---

Feel free to modify or add more details as needed!