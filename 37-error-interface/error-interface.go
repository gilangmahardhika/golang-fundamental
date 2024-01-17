package main

import (
	"errors"
	"fmt"
)

func devidedByZero(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("can't divide by zero")
	} else {
		return value / divider, nil
	}
}

func main() {
	devider := devidedByZero
	result, err := devider(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}
