package main

import (
	"fmt"
)

func PromptUser() [10]int {
	fmt.Println("Please enter up to 10 integers")
	inputs := [10]int{}
	count := 0

	for count != 10 {
		var userInput int
		_, err := fmt.Scanf("%d", &userInput)
		if err != nil {
			fmt.Println("You did not enter an integer, that means you want to stop here")
			break
		}
		inputs[count] = userInput
		count++
	}

	return inputs
}
func BubbleSort(toSort []int) {
	var swapped bool = true
	for swapped {
		var currentIndex int = 0
		swapped = false
		for currentIndex != len(toSort)-1 {
			if toSort[currentIndex] > toSort[currentIndex+1] {
				Swap(toSort, currentIndex)
				swapped = true
			}
			currentIndex++
		}
	}
}

func Swap(toSort []int, currentIndex int) {
	toSort[currentIndex], toSort[currentIndex+1] = toSort[currentIndex+1], toSort[currentIndex]
}

func main() {
	arrayToSort := PromptUser()
	BubbleSort(arrayToSort[:])
	fmt.Println(arrayToSort)

}
