package main

import (
	"TypeChecker/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	st1, _ := ioutil.ReadFile("Input/ST2")
	in1, _ := ioutil.ReadFile("Input/IN2")

	var ist1 interface{}
	var iin1 interface{}

	_ = json.Unmarshal(st1, &ist1)
	_ = json.Unmarshal(in1, &iin1)

	fmt.Println(Utils.Validate(iin1, ist1))

}
