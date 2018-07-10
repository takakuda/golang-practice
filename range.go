package main

import "fmt"

func main() {
	var arr [4]string

	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"
	arr[3] = "d"

	for i, s := range arr {
		fmt.Println(i, s)
	}
}
