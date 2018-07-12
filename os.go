package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ファイルを開く
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// プログラムが終了したらファイルを閉じる
	defer file.Close()

	// 12byte格納可能なスライスを用意する
	message := make([]byte, 12)

	// ファイル内のデータをスライスに読み出す
	_, err = file.Read(message)
	if err != nil {
		log.Fatal(err)
	}

	// 文字列にして表示
	fmt.Print(string(message))
}
