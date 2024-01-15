package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Gilang"
	names[1] = "Putera"
	names[2] = "Mahardhika"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		100,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
