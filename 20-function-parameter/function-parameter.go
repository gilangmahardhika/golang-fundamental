package main

import "fmt"

func sayMyAge(age int8, name string) {
	fmt.Println("Hi, my name is", name, "I am", age, "years old")
}

func main() {
	sayMyAge(36, "Gilang")
}
