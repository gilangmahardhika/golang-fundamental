package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi() {
	fmt.Println("Hello, my name is", customer.Name)
}

func (customer Customer) sayOK(name string) {
	fmt.Println("Ok,", name, "From", customer.Name)
}

func main() {
	var name = Customer{"Gilang", "Jakarta", 36}
	name.sayHi()
	name.sayOK("Whoever")
}
