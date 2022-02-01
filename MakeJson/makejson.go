package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter an address: ")
	scanner.Scan()
	address := scanner.Text()

	myMap := map[string]string{
		"name":    name,
		"address": address,
	}

	barr, err := json.Marshal(myMap)
	if err == nil {
		fmt.Println(string(barr))
	}

}
