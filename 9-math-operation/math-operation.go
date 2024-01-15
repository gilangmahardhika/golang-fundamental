package main

import "fmt"

func main() {
	var a int32
	var b int32
	var c int32
	a = 10
	b = 20
	c = a + b
	fmt.Println((c))

	a += 59
	fmt.Println(a)

	c--
	fmt.Println(c)
}
