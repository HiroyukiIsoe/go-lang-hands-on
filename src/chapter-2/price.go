package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type a price")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	p := float64(n)
	fmt.Println(int(p * 1.1))
}

// パッケージ登録サボったのでinput関数を個別に定義
func input(meg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(meg + ": ")
	scanner.Scan()
	return scanner.Text()
}
