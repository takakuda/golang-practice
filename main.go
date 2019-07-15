package main

import (
  "fmt"
)

func main() {
  s := make([]int, 0, 0)
  fmt.Printf("(A) len=%d, cap=%d\n", len(s), cap(s))

  s = append(s, 1)
  fmt.Printf("(B) len=%d, cap=%d\n", len(s), cap(s))

  s = append(s, []int{1, 2, 3, 4}...)
  fmt.Printf("(C) len=%d, cap=%d\n", len(s), cap(s))

  s = append(s, 5)
  fmt.Printf("(D) len=%d, cap=%d\n", len(s), cap(s))

  s = append(s, 6, 7, 8, 9)
  fmt.Printf("(E) len=%d, cap=%d\n", len(s), cap(s))
}
