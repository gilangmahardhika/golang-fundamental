package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Gilang")
	fmt.Println(slice3)
	slice3[1] = "Edited Data"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Gilang"
	newSlice[1] = "Mahardhika"

	fmt.Println(newSlice)

}
