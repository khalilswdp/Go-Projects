package main

import (
	"fmt"
	"math"
)

func main() {
	acc, vel, dis := PromptUserEquation()
	var computeDisplacement = GenDisplaceFn(acc, vel, dis)

	// We can make a loop just asking about time here...
	var time = funcName("Time")
	fmt.Print("The displacement is: ")
	fmt.Println(computeDisplacement(time))

}

func GenDisplaceFn(acc float64, vel float64, dis float64) func(float64) float64 {
	return func(time float64) float64 {
		return (1/2)*acc*math.Pow(time, 2) + vel*time + dis
	}
}

func PromptUserEquation() (float64, float64, float64) {
	var accInp = funcName("Acceleration")
	var velInp = funcName("Initial Velocity")
	var disInp = funcName("Initial Displacement")

	return accInp, velInp, disInp
}

func funcName(inputName string) float64 {
	fmt.Printf("Enter %s: ", inputName)
	var userInput float64
	_, errAcc := fmt.Scanf("%f", &userInput)
	if errAcc != nil {
		fmt.Println("You did not enter a valid value, replacing it with 0.0")
		userInput = 0.0
	}
	return userInput
}
