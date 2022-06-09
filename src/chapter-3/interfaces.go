package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Data is interface for Mydata.
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

// Initial is init method
func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

// PrintData is println all data.
func (md *Mydata) PrintData() {
	if md != nil {
		fmt.Println("Name: ", md.Name)
		fmt.Println("Data: ", md.Data)
	} else {
		fmt.Println("** This is Nil value.**")
	}
}

// Yourdata is structure
type Yourdata struct {
	Name string
	Mail string
	Age int
}

func (yd *Yourdata) SetValue(vals map[string]string) {
	yd.Name = vals["name"]
	yd.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	yd.Age = n
}

func (yd *Yourdata) PrintData() {
	fmt.Printf("I'm %s. (%d).\n", yd.Name, yd.Age)
	fmt.Printf("mail: %s.\n", yd.Mail)
}

func main()  {
	var nilObj *Mydata
	nilObj.PrintData()
	ob := [2]Data{}
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})
	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name": "Mami",
		"mail": "mami@mume.mo",
		"age": "34",
	})

	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}
