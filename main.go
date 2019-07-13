package main

import (
  "fmt"
)

func main() {
LOOP:
  for {
    for {
      for {
        fmt.Println("開始")
        break LOOP
      }
      fmt.Println("通らない")
    }
    fmt.Println("通らない")
  }
  fmt.Println("完了")
}
