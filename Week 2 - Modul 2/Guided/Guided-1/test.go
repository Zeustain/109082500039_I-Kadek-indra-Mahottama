package main

import "fmt"

func main() {
	var satu, dua, tiga string
	var temp string

	fmt.Print("Masukan Input: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan Input: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan Input: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output Awal: " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output Akhir: " + satu + " " + dua + " " + tiga)
}
