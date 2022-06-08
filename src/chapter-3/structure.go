package main

import "fmt"

// Myadata is structure
type Mydata struct {
	Name string
	Data []int
}

// PrintData is println all data.
func (md Mydata) PrintData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** end ***")
}

func main()  {
	taro := Mydata{"Taro", []int{10, 20, 30},}
	hanako := Mydata{
		Name: "Hanako",
		Data: []int{90, 80, 70},
	}
	fmt.Println(taro)
	fmt.Println(hanako)
	rev(&taro)
	fmt.Println(taro)
	taro.PrintData()
}

func rev(md *Mydata) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) -1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
