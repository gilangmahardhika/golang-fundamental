package main

import "fmt"

func logging() {
	fmt.Println("Finished")
}

func runApp(value int) {
	defer logging()
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println("Result", result)
	logging()
}

func main() {
	runApp(0)
}
