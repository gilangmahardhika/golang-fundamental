package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	sayHello := getHello
	result := sayHello("Gilang")
	fmt.Println(result)
}
