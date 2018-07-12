package main

import (
	"log"
	"os"
)

func main() {
	// ファイルを生成
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// プログラムが終わったらファイルを閉じる
	defer file.Close()
}
