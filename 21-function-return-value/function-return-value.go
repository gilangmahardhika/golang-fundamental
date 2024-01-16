package main

import "fmt"

func myName(name string) string {
	return name
}

type Person struct {
	Name, Title string
}

func myData(name string, title string) Person {
	var p Person = Person{name, title}
	return p
}

func main() {
	fmt.Println("my name is", myName("Gilang"))
	fmt.Println(myData("Gilang", "Programeer"))
}
