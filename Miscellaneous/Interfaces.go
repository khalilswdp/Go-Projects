package main

import "fmt"

type Speaker interface {
	Speak()
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (c Cat) Speak() {
	fmt.Println(c.name)
}

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var s1, s2 Speaker
	var c1 = Cat{"Brian"}
	var d1 *Dog
	s1 = &c1
	s2 = d1
	s1.Speak()
	// Dynamic
	if s2 != nil {
		s2.Speak()
	}
}
