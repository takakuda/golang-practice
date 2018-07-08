package main

import (
	"errors"
	"fmt"
	"log"
)

func div(i, j int) (result int, err error) {
	if j == 0 {
		err = error.New("devied by zero")
		return // return 0, errと同じ
	}
	result = i / j
	return
}

func main() {
	n, err := div(10, 0)
	if err != nil {
		// エラーを出力しプログラムを終了する
		log.Fatal(err)
	}
	fmt.Println(n)
}
