package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	close(ch)
	i, ok := <-ch
	fmt.Println(i, ok)
}
