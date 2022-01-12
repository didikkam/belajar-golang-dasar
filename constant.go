package main

import "fmt"

func main() {
	const firstName string = "Didik"
	const lastName = "Mukmin"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		first = "Didik"
		last  = "Mukmin"
	)
}
