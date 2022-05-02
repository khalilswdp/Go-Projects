package main

import "fmt"

type MyInt int

func (mi MyInt) Double() int {
	return int(mi * 2)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())
}
