package main

import "fmt"

func main() {
	var name string
	name = "Gilang Mahardhika"

	if name == "Gilang" {
		fmt.Println(name)
	} else {
		fmt.Println("Who are you?")
	}

	// Short statement
	if lastName := "Mahardhika"; lastName == "Mahardhika" {
		fmt.Println(lastName)
	}
}
