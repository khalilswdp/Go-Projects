package main

import "fmt"

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	fmt.Println(getMax(1, 2, 3, -50, 500))
}
