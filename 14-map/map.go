package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Gilang",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["lastName"] = "Mahardhika"
	fmt.Println(person["lastName"])

	fmt.Println(len(person))

	delete(person, "address")
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "Learning Golang"

	fmt.Println(book)

}
