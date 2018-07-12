package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func main() {
	task := New.Task(1, "buy the milk")
	fmt.Printf("#s", task)
}
