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
  */
  fmt.Println("My", "name", "is", "golang")
  fmt.Printf("数値=%d\n", 5)
  fmt.Printf("10進法=%d 2進法=%b 8進法=%o 16進法=%x\n", 17, 17, 17, 17)
  fmt.Printf("%d年%d月%d日\n", 2015, 12, 1)
}
