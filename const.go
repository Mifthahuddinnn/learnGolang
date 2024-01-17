package main

import "fmt"

func main() {
	const firstName = "Mifthah"
	const lastName = "Huddin"
	const age = 19
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	const (
		Universitas = "Mulawarman"
		Hoby        = "Main Game"
	)

	fmt.Println(Universitas)
	fmt.Println(Hoby)
}
