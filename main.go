package main

import (
  "fmt"
)

func main() {
  var x interface{} = 1
  switch x.(type) {
  case bool:
    fmt.Println("bool")
  case int, uint:
    fmt.Println("integer or unsigned integer")
  default:
    fmt.Println("don`t know")
  }
}
