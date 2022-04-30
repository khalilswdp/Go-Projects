package main

import "fmt"

var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func applyIt(funcArg func(int) int, para int) int {
	return funcArg(para)
}

func main() {
	funcVar = incFn
	fmt.Println(funcVar(1))

	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(func(x int) int { return x + 3 }, 2))

}
