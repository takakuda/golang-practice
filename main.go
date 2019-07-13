package main

import (
  "fmt"
)

func later() func(string) string {
  var store string
  return func(next string) string {
    s := store
    store = next
    return s
  }
}

func main() {
  f := later()

  fmt.Println(f("Golsng"))
  fmt.Println(f("is"))
  fmt.Println(f("awesome!"))
}
