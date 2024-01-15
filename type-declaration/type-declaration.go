package main

import "fmt"

func main() {
	type PassportNumber string

	var passportNumber PassportNumber = "X70000002"

	fmt.Println("No Psspornya", passportNumber)
}
