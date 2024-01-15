package main

import "fmt"

func main() {
	var lastPoint int8 = 90
	var presence int8 = 75

	var graduatePoint bool = lastPoint > 85
	var graduatePresence bool = presence > 50

	var graduate bool = graduatePoint && graduatePresence

	fmt.Println(graduate)
}
