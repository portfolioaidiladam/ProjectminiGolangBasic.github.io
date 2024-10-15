package main

import "fmt"

func main() {

	type NoKTP string

	var ktpAidil NoKTP = "1111111111"

	var contoh string = "222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpAidil)
	fmt.Println(contohKtp)

}
