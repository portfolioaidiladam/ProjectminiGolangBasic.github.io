package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("Aidil"))
	fmt.Println(misal("Aidil"))
}
