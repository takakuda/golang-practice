package main

import (
  "fmt"

  _ "./animals"
)

func main() {
  /*
  fmt.Println(AppName())

  fmt.Println(animals.ElephantFeed())
  fmt.Println(animals.MonkeyFeed())
  fmt.Println(animals.RabbitFeed())
  fmt.Println("My", "name", "is", "golang")
  fmt.Printf("数値=%d\n", 5)
  fmt.Printf("10進法=%d 2進法=%b 8進法=%o 16進法=%x\n", 17, 17, 17, 17)
  fmt.Printf("%d年%d月%d日\n", 2015, 12, 1)
  fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
  fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})
  fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})
  */
  n := one()
  fmt.Println(n)
}

func one() int {
  return 1
}
