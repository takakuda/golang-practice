package main

import (
  "fmt"
)

func main() {
  if x, y := 1, 2; x < y {
    fmt.Printf("x(%d) is less than y(%d)\n", x, y)
  }
}
