package main

import "fmt"

func main() {
	var name string
	name = "Anton"

	switch name {
	case "Gilang":
		fmt.Println("Namanya", name)
	default:
		fmt.Println("Who are you?")
	}

	switch lastName := "Mahardhika"; lastName == "Mahardika" {
	case true:
		fmt.Println("Correct name")
	case false:
		fmt.Println("Incorrect name")
	}

	var length int
	length = len(name)
	switch {
	case length == 5:
		fmt.Println("True")
	case length <= 5:
		fmt.Println("Too short")
	default:
		fmt.Println("Too long")
	}
}
