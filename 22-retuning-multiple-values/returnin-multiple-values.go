package main

import "fmt"

type Bio struct {
	Name, Title string
	Age         int8
}

func biodata(name string, title string, age int8) (string, string, int8) {
	bio := Bio{name, title, age}
	return bio.Name, bio.Title, bio.Age
}

func main() {
	name, title, age := biodata("Gilang", "Programmer", 36)
	fmt.Println("My name is", name, "I am a ", title, "my age is", age)

	otherName, _, _ := biodata("gilang", "aaa", 21)
	fmt.Println(otherName)
}
