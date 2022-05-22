package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Print("123 * 45 = ")
	// fmt.Println(123*45)
	name := input("type your name")
	fmt.Println("Hello, " + name + "!!")
}

func input(meg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(meg + ": ")
	scanner.Scan()
	return scanner.Text()
}
