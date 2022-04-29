package main

import "fmt"

func pbr(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}
func main() {
	a := [3]int{1, 2, 3}
	pbr(&a)
	fmt.Println(a)
}
