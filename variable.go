package main

import "fmt"

func main() {
	var name string

	name = "Didik Abdul"
	fmt.Println(name)

	name = "Didik Mukmin"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age int8 = 24
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	country = "America"
	fmt.Println(country)

	var (
		firstName = "Didik"
		lastName  = "Mukmin"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
