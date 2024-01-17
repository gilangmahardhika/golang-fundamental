package main

import "fmt"

func endApp() {
	fmt.Println("Finished")
	message := recover()
	fmt.Println(message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}

	fmt.Println("Running")

}

func main() {
	runApp(true)
}
