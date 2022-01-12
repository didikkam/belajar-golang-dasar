package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "Didik",
		"address": "Tuban",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Didik"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")

	fmt.Println(book)
}
