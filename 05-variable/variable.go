package main

import "fmt"

func main() {
	var name string

	name = "Gilang Putera"
	fmt.Println(name)

	name = "Gilang Mahardhika"
	fmt.Println(name)

	var lastName = "Mahardhika"
	fmt.Println(lastName)

	var age int8 = 30
	fmt.Println(age)

	middleName := "Putera"
	fmt.Println(middleName)

	var (
		firstName = "Gilang"
		lastName2 = "Mahardhika"
	)

	fmt.Println(firstName)
	fmt.Println(lastName2)

}
