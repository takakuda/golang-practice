package main

import (
  "fmt"
)

func main() {
  a := plus(1, 2)
  fmt.Printf("%v", a)
}

func plus(x, y int) int {
  return x + y
}
