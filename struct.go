package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var aidil Customer
	fmt.Println(aidil)

	aidil.Name = "Aidil Adam"
	aidil.Address = "Indonesia"
	aidil.Age = 30
	fmt.Println(aidil)
	fmt.Println(aidil.Name)
	fmt.Println(aidil.Address)
	fmt.Println(aidil.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	// method struct

	budi.sayHello("Agus")
	aidil.sayHello("Agus")
	joko.sayHello("Agus")
}
