package main

import (
  "fmt"
)

func main() {
  switch n := 2; n {
  case 1, 3, 5, 7, 9:
    fmt.Printf("%d is odd\n", n)
  case 2, 4, 6, 8, 10:
    fmt.Printf("%d is even\n", n)
  }
}
