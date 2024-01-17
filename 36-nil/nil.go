package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "Gilang" {
		return map[string]string{
			"name": name,
		}
	} else {
		return nil
	}
}

func main() {
	fmt.Println(newMap("Gilang"))
}
