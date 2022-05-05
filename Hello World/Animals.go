package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}
func (cow Cow) Move() {
	fmt.Println("walk")
}
func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("worms")
}
func (bird Bird) Move() {
	fmt.Println("fly")
}
func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}
func (snake Snake) Move() {
	fmt.Println("slither")
}
func (snake Snake) Speak() {
	fmt.Println("hsss")
}

var animalMap = make(map[string]Animal)

func main() {
	// “cow”, “bird”, or “snake”

	for {
		fmt.Print("> ")
		command := ""
		name := ""
		animalOrAction := ""
		_, err := fmt.Scanln(&command, &name, &animalOrAction)
		if err != nil {
			return
		}

		switch command {
		case "newanimal":

			var newAnimal Animal

			switch animalOrAction {
			case "cow":
				newAnimal = Cow{}
			case "bird":
				newAnimal = Bird{}
			case "snake":
				newAnimal = Snake{}
			default:
				fmt.Println("The input is wrong. Try again")
				continue
			}

			animalMap[name] = newAnimal
			fmt.Println("Created it!")

		case "query":
			var animalInstance = animalMap[name]
			switch animalOrAction {
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
		default:
			fmt.Println("The Command is wrong!")

		}
	}
}
