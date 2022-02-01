package main

import (
	"fmt"
	"log"
)

var num float64

func main() {
	fmt.Println("Enter a Float number=")
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Printf("[Error] Invalid input !")
	}

	fmt.Printf("The number we have entered is '%v'.\n", num)
	fmt.Printf("Truncated version of '%v' is '%v'.\n", num, int64(num))
}
