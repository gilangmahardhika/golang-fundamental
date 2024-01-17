package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "mr. " + man.Name
	fmt.Println("on method", man.Name)
}

func main() {
	man := Man{"Gilang"}
	man.Married()

	fmt.Println(man)
}
