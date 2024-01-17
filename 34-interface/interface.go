package main

import "fmt"

type HasName interface {
	getName() string
}

type HasAge interface {
	getAge() int
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

func getAge(hasAge HasAge) {
	fmt.Println("Age", hasAge.getAge())
}

type Person struct {
	Name string
}

type Customer struct {
	Name string
	Age  int
}

func (person Person) getName() string {
	return person.Name
}

func (customer Customer) getName() string {
	return customer.Name
}
func (customer Customer) getAge() int {
	return customer.Age
}

func main() {
	var person Person
	person.Name = "Gilang"

	sayHello(person)

	var customer Customer
	customer.Age = 36
	getAge(customer)

}
