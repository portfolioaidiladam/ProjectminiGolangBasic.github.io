package main

import "fmt"

func main() {
	const (
		firstName string = "Aidil"
		lastName         = "Adam"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}
