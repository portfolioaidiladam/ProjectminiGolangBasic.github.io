package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Aidil"
	//person["address"] = "Subang"

	person := map[string]string{
		"name":    "Aidil",
		"address": "Cikarang Selatan",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Aidil"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
