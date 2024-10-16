package main

import "fmt"

func getFullName() (string, string) {
	return "Aidil", "Adam"
}

func main() {
	/* firstName, lastName := getFullName()
	fmt.Println(firstName, lastName) */

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
