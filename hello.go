package main

import "fmt"

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	x, y := 3, 4
	x, _ = swap(x, y)
	fmt.Println(x)
}
