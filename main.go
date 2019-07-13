package main

import (
  "fmt"
)

func main() {
  var x interface{} = 1
  if x == nil {
    fmt.Println("x is nil")
  } else if i, isInt := x.(int); isInt {
    fmt.Println("x is integer : %d\n", i)
  } else if s, isString := x.(string); isString {
    fmt.Println(s)
  } else {
    fmt.Println("unsupported type!")
  }
}
