package main

import (
	"fmt"
	"math"
)

func main() {
	acc, vel, dis := PromptUserEquation()
	var computeDisplacement = GenDisplaceFn(acc, vel, dis)
	time := PromptUserTime()

	fmt.Print("The displacement is: ")
	fmt.Println(computeDisplacement(time))

}

func PromptUserTime() float64 {
	fmt.Print("Enter Time: ")
	var timeInp float64
	_, errAcc := fmt.Scanf("%f", &timeInp)
	if errAcc != nil {
		fmt.Println("You did not enter a valid value, replacing it with 0.0")
		timeInp = 0.0
	}

	return timeInp
}

func GenDisplaceFn(acc float64, vel float64, dis float64) func(float64) float64 {
	return func(time float64) float64 {
		return (1/2)*acc*math.Pow(time, 2) + vel*time + dis
	}
}

func PromptUserEquation() (float64, float64, float64) {
	fmt.Print("Enter Acceleration: ")
	var accInp float64
	_, errAcc := fmt.Scanf("%f", &accInp)
	if errAcc != nil {
		fmt.Println("You did not enter a valid value, replacing it with 0.0")
		accInp = 0.0
	}

	fmt.Print("Enter Initial Velocity: ")
	var velInp float64
	_, errVel := fmt.Scanf("%f", &velInp)
	if errVel != nil {
		fmt.Println("You did not enter a valid value, replacing it with 0.0")
		velInp = 0.0
	}

	fmt.Print("Enter Initial displacement: ")
	var disInp float64
	_, errDis := fmt.Scanf("%f", &disInp)
	if errDis != nil {
		fmt.Println("You did not enter a valid value, replacing it with 0.0")
		disInp = 0.0
	}

	return accInp, velInp, disInp
}
