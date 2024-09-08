package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var stringsList []string

	for {
		fmt.Print("String to add: ")
		fmt.Scanln(&input)
		stringsList = append(stringsList, input)

		var continueInput string
		fmt.Print("Continue? [Y/n]: ")
		fmt.Scanln(&continueInput)

		if strings.ToLower(continueInput) == "n" {
			break
		}
	}

	output := strings.Join(stringsList, ", ")
	fmt.Println(output)
}
