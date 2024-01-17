package main

import "fmt"

type Filter func(string, int) (string, int)

func sayHelloWithFilter(name string, age int, filter Filter) {
	nameFiltered, age := filter(name, age)
	fmt.Println("Hello", nameFiltered, "I am", age)

}

func spamFilter(name string, age int) (string, int) {
	if name == "Fuck" {
		return "filtered", age
	} else {
		return name, age
	}
}

func main() {
	sayHelloWithFilter("Fuck", 36, spamFilter)
}
