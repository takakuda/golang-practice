package main

func callByValue(i int) {
	i = 20
}

func callByRef(i *int) {
	*i = 20
}

func main() {
	file, err := os.Open("./error.go")
	if err != nil {
		// エラー処理
	}
	// 関数を抜ける前に必ず実行される
	defer file.Close()
	// 正常処理
}
