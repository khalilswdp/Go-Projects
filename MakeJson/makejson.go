package main

import (
	"encoding/json"
	"fmt"
)

var name string
var address string

func main() {

	m := make(map[string]string)

	fmt.Print("Enter a name= ")
	fmt.Scanln(&name)

	fmt.Print("Enter an address= ")
	fmt.Scanln(&address)

	m["name"] = name
	m["address"] = address

	output, _ := json.Marshal(m)

	fmt.Printf(string(output))
}
