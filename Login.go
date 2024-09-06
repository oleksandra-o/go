package main

import (
	"fmt"
)

const (
	maxAttempts = 3
	loginID     = "admin"
	password    = "Pa$$w0rd"
)

func main() {
	attempts := 0
	for attempts < maxAttempts {
		var inputID, inputPwd string
		fmt.Print("Enter loginID: ")
		fmt.Scanln(&inputID)
		fmt.Print("Enter password: ")
		fmt.Scanln(&inputPwd)

		if inputID == loginID && inputPwd == password {
			fmt.Println("Login Successful!")
			return
		} else if inputID == loginID && inputPwd != password {
			fmt.Println("Password Incorrect!")
		} else if inputID != loginID && inputPwd == password {
			fmt.Println("Login Incorrect!")
		} else {
			fmt.Printf("Try again - %d attempts left!\n", maxAttempts-attempts-1)
		}

		attempts++
	}

	fmt.Println("Account locked - Contact 800-123-4567")
}
