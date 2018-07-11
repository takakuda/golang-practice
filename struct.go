package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	var task Task = Task{
		1, "buy the milk", true,
	}
	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)
}
