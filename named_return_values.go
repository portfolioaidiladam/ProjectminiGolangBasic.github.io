package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Aidil"
	middleName = "Adam"
	lastName = "Baik"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
