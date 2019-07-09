package main

import (
  "fmt"
)

func plus(x, y int) int {
  return x + y
}

var plusAlias = plus

func main() {
  fmt.Println(plusAlias(10, 5))
}
