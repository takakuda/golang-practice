package main

import (
  "fmt"
)

func main() {
  a := [5]int{1, 2, 3, 4, 5}
  s := a[0:2]
  fmt.Println(s)
}
