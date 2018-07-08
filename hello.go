package main

import "fmt"

func main() {
	a, b := 10, 100
	if a > b {
		fmt.Println("a is lager than b")
	} else if a < b {
		fmt.Println("b is lager than a")
	} else {
		fmt.Println("a equal b")
	}
}
