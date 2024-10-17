package helper

import "fmt"

var version = "1.0.0"
var Application = "golang"

// tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Aidil")
	fmt.Println(version)
}
