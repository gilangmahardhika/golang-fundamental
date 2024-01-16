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
}
