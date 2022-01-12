package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Didik"
	names[1] = "Abdul"
	names[2] = "Mukmin"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
}
