package main

import "fmt"

func pbc(x [3]int) int {
	return x[0]
}
func main() {
	a := [3]int{1, 2, 3}
	fmt.Print(pbc(a))
}
