package main

import (
	"fmt"
)

func main() {
	s1 := []string{"a", "b"}
	s2 := []string{"c", "d"}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
