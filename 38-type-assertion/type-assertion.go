package main

import "fmt"

func random() interface{} {
	return "Whoa"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	result2 := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println(result2, "Unknown")
	}
}
