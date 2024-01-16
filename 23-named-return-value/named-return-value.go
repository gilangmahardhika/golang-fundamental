package main

import "fmt"

type Bio struct {
	Name, Title string
	Age         int8
}

func biodata(name string, title string, age int8) (fullName string, job string, yearsOld int8) {
	bio := Bio{name, title, age}
	fullName = bio.Name
	job = bio.Title
	yearsOld = bio.Age
	return
}

func main() {
	name, title, age := biodata("Gilang", "Programmer", 36)
	fmt.Println("My name is", name, "I am a ", title, "my age is", age)

	otherName, _, _ := biodata("gilang", "aaa", 21)
	fmt.Println(otherName)
}
