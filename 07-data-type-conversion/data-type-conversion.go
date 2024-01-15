package main

import "fmt"

func main() {
	var value32 int32 = 100
	var value64 = int64(value32)

	fmt.Println(value32)
	fmt.Println(value64)

	var name string = "Gilang"
	var e = name[0]
	var character string = string(e)
	fmt.Println(character)
}
