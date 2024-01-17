package main

import "fmt"

type EmptyInterface interface{}

func empty(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return "dua"
	} else {
		return false
	}
}

func main() {
	var empty interface{} = empty(3)
	fmt.Println(empty)
}
