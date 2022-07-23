package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Mydata is json structure.
type Mydata struct {
	Name string
	Mail string
	Tel string
}

// Str get string value.
func (m *Mydata) Str() string {
	return "<\"" + m.Name + "\" " + m.Mail + ", " + m.Tel + ">"
}

func main() {
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	// var data []interface{}
	// er = json.Unmarshal(s, &data)
	var items []Mydata
	er = json.Unmarshal(s, &items)
	if er != nil {
		panic(er)
	}

	// for i, im := range data {
	// 	m := im.(map[string]interface{})
	// 	fmt.Println(i, m["name"].(string), m["mail"].(string), m["tel"].(string))
	// }

	for i, im := range items {
		println(i, im.Str())
	}
}
