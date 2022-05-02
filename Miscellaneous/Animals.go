package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) init(f, l, n string) {
	animal.food = f
	animal.locomotion = l
	animal.noise = n
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	// “cow”, “bird”, or “snake”

	var cow, bird, snake Animal

	cow.init("grass", "walk", "moo")
	bird.init("worms", "fly", "peep")
	snake.init("mice", "slither", "hsss")

	for {
		fmt.Print("> ")
		animal := ""
		action := ""
		_, err := fmt.Scanln(&animal, &action)
		if err != nil {
			return
		}
		
		var animalInstance Animal

		switch animal {
		case "cow":
			animalInstance = cow
		case "bird":
			animalInstance = bird
		case "snake":
			animalInstance = snake
		default:
			fmt.Println("The input is wrong. Try again")
			continue
		}

		switch action {
		case "eat":
			animalInstance.Eat()
		case "move":
			animalInstance.Move()
		case "speak":
			animalInstance.Speak()
		default:
			fmt.Println("The input is wrong. Try again")
			continue
		}
	}
}
