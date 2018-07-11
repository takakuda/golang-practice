package main

import "fmt"
import "log"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()

	a := []int{1, 2, 3}
	fmt.Println(a[10])
}
