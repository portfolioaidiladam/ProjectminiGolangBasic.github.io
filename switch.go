package main

import "fmt"

func main() {
	name := "Aidil"

	switch name {
	case "Aidil":
		fmt.Println("Hello Aidil")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	name = "AdamAdamAdam"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
