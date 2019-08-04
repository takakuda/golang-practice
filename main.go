package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	delete(m, 2)
	fmt.Println(len(m))
}
