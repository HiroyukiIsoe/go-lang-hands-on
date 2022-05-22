package hello

import (
	"bufio"
	"fmt"
	"os"
)

func Input(meg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(meg + ": ")
	scanner.Scan()
	return scanner.Text()
}
