package main

import "fmt"

var roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(input string) int {
	// Read Block that goes together
	// Calculate block value
	var res int
	if len(input) == 0 {
		return 0
	}
	var currVal = roman[input[0]]
	if len(input) == 1 {
		return currVal
	}
	if currVal >= roman[input[1]] {
		res = currVal
	} else {
		res = -currVal
	}
	return res + romanToInt(input[1:])
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
