package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Gilang"
	customer.Address = "Jakarta"
	customer.Age = 36

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	customer2 := Customer{"Mahardhika", "Jakarta", 36}
	fmt.Println(customer2)

	customer3 := Customer{
		Name:    "Putera",
		Address: "Yogyakarta",
		Age:     36,
	}
	fmt.Println(customer3)
}
