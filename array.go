package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "Mifthahuddin"
	names[1] = "King"
	names[2] = "Udin"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var nomor = [3]int{
		90,
		10,
		30,
	}
	fmt.Println(nomor[0])

	var number = [3]int{
		19,
		21,
		123,
	}
	fmt.Println(number[1])

}
