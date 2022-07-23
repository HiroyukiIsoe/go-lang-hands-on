package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// write text function
	wt := func (f *os.File, s string)  {
		_, er := f.WriteString(s + "\n")
		if er != nil {
			panic(er)
		}
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}
	// defer close.
	defer f.Close()

	fmt.Println("*** start ***")
	wt(f, "*** start ***")
	for {
		a := input("type message")
		if a == "" {
			break
		}
		wt(f, a)
	}
	wt(f, "*** end ***\n\n")
	fmt.Println("*** end ***")
}

// パッケージ登録サボったのでinput関数を個別に定義
func input(meg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(meg + ": ")
	scanner.Scan()
	return scanner.Text()
}
