package main

import "fmt"

func main() {
	n := 0
	for {
		n++
		if n > 15 {
			break
		}
		switch {
		case n%15 == 0:
			fmt.Println("FizzBuzz")
		case n%5 == 0:
			fmt.Println("Fizz")
		case n%3 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(n)
		}
	}
}
