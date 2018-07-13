package main

import (
	"log"
	"os"
)

func main() {
	person := &Person{
		ID:     1,
		Name:   "Gopher",
		Email:  "gopher@example.org",
		Age:    5,
		Adress: "",
		memo:   "golang lover",
	}

	// ファイルを開く
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// エンコーダの取得
	encoder := json.NewEncoder(file)

	// jsonエンコーダーしたデータの書き込み
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
}
