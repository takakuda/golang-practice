package main

import "fmt"

var month map[int]string = map[int]string{}

func main() {
	month := map[int]string{
		1: "january",
		2: "February",
	}
	_, ok := month[1]
	if ok {
		fmt.Println(month)
	}
}
