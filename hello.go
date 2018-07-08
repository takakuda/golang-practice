package main

import "fmt"

func main() {
	file, err := os.Open("hello.go")
	if err != nil {
		// エラー処置
		// returnなどで処理を別の場所に抜ける
	}
	// fileを用いた処理
}
