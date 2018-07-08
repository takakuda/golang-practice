package main

import "fmt"

func main() {
	n := 0
	for {
		n++
		if n > 15 {
			break
		}
		switch n {
		case 15:
			fmt.Println("FizzBuzz")
		case 5, 10:
			fmt.Println("Buzz")
		case 3, 6, 9, 12:
			fmt.Println("Fizz")
		default:
			fmt.Println(n)
		}
	}
}
