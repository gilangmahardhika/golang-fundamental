package main

import "fmt"

func main() {
	var counter int8 = 0

	fmt.Println(counter)

	increment := func() {
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}
