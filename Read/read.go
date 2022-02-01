package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type name struct {
	fname string
	lname string
}

func main() {

	// Read the name of the text file
	var person string

	fmt.Print("Enter the the name of the text file: ")
	fmt.Scanf("%s", &person)
	fmt.Println("The name of the text file is: ", person)

	// Open the file

	f, err := os.Open("./" + person)
	check(err)
	scanner := bufio.NewScanner(f)

	// Read the file one line at a time

	scanner.Split(bufio.ScanLines)
	var text string

	slice := make([]name, 0, 3)

	for scanner.Scan() {
		text = scanner.Text()
		words := strings.Fields(text)

		if len(words) != 2 {
			fmt.Printf("Beware, there's a line that doesn't satisfy the dependencies")
		} else {
			//
			newName := name{fname: words[0], lname: words[1]}
			slice = append(slice, newName)
		}
	}

	for _, nameStruct := range slice {
		fmt.Printf("First Name: %s //--// Last Name: %s\n", nameStruct.fname, nameStruct.lname)

		fmt.Println("")
	}
	f.Close()
}
