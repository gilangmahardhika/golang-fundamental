package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Sumedang"

	*address2 = Address{"Jogja", "DIY", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
