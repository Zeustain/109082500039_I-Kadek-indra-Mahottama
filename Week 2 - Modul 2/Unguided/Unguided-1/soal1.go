package main

import "fmt"

func main() {
	var a int

	fmt.Println("Masukan tahun: ")
	fmt.Scanln(&a)

	if a%400 == 0 {
		fmt.Printf("Tahun: %.d\nKabisat: True", a)
	} else if a%4 == 0 && a%100 != 0 {
		fmt.Printf("Tahun: %.d\nKabisat: True", a)
	} else {
		fmt.Printf("Tahun: %.d\nKabisat: False", a)
	}
}
