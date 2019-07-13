package main

import (
  "fmt"
)

func main() {
  for i := 0; i < 100; i++ {
    if (i % 2 == 1) {
      continue
    }
    fmt.Println(i)
    i++
  }
  fmt.Println(16 % 2 == 1)
}
