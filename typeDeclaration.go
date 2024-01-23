package main

import "fmt"

func main() {

	type NoKTP string
	var eKTP NoKTP = "123123132"
	fmt.Println(eKTP)

	var contoh string = "231232132"
	var contohKTP NoKTP = NoKTP(contoh)
	fmt.Println(contohKTP)

}
