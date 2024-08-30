package main

import (
	"fmt"
	taskresulterror "github.com/zukaChachava/task/task/single/task-result-error"
)

func main() {

	result := taskresulterror.NewTask[*Person, error](func() (*Person, error) {
		return &Person{name: "Zura", lastname: "Chachava"}, nil
	}).Run()

	value, err := result.Wait()

	if err != nil {
		panic(err)
	}

	fmt.Println(*value)
}

type Person struct {
	name     string
	lastname string
}
