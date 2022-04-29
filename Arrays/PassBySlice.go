package main

import "fmt"

func pbs(s []int) {
	s[1] = s[1] + 1
}
func main() {
	a := []int{1, 2, 3}
	pbs(a)
	fmt.Println(a)
}
