package main

import "fmt"

func main() {
	var counter int8
	counter = 1

	for counter <= 12 {
		fmt.Println(counter)
		counter++
	}

	for looper := 1; looper <= 13; looper++ {
		fmt.Println(looper)
	}

	slice := []string{"Gilang", "Putera", "Mahardhika"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, name := range slice {
		fmt.Println("name", "=", name)
	}

	person := make(map[string]string)
	person["name"] = "Gilang"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, value)
	}
}
