package main

import "fmt"

func main() {
	type NoKTP string

	var noKtpDidik NoKTP = "1234567890123456"
	fmt.Println(noKtpDidik)
}
