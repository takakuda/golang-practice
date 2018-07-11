package main

import "fmt"
import "log"

func main() {
  a := int{1, 2, 3}
  for i := 0; < 10; ++ {
    if i >=len(a) {
      panic(errors.New("index out of range"))
    }
    fmt.Println(a[i])
  }
