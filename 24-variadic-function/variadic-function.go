package main

import "fmt"

func sumAll(numbers ...int) (total int) {
	total = 0
	for _, value := range numbers {
		total += value
	}
	return
}

func main() {
	numbers := []int{10, 11, 12}
	total := sumAll(numbers...)
	fmt.Println(total)

}
