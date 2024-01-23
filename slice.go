package main

import "fmt"

func main() {
	names := [...]string{"Mifthah", "Huddin", "Metropolis", "Batman", "Gotham"}
	slice := names[3:5]
	fmt.Println(slice[0])
	slice2 := names[:2]
	fmt.Println(slice2)
	slice3 := names[2:]
	fmt.Println(slice3)
	slice5 := names[:]
	fmt.Println(slice5)

	fmt.Println(len(slice))
	fmt.Println(cap(slice5))

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur baru")
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
}
