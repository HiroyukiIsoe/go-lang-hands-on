package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read text function.
	rt := func (f *os.File)  {
		// s, er := ioutil.ReadAll(f)
		// if er != nil {
		// 	panic(er)
		// }
		// fmt.Println(string(s))
		r := bufio.NewReaderSize(f, 4096)
		for i := 0; true; i++ {
			s, _, er := r.ReadLine()
			if er != nil {
				break
			}
			fmt.Println(i, ":", string(s))
		}
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}
	// defer close.
	defer f.Close()

	fmt.Println("<< start >>")
	rt(f)
	fmt.Println("<< end >>")
}
