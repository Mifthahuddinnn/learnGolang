package main

import "fmt"

func main() {
	nilai := 90
	absensi := 12

	var nilaiMinimum = nilai >= 75
	var absensiMinimum = absensi >= 12

	var kelulusan = nilaiMinimum && absensiMinimum
	fmt.Println(kelulusan)

}
