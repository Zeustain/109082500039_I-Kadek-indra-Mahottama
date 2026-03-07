package main

import "fmt"

func main() {
	var gram, kg int
	var sisakg, sisaGram int
	var ttlkg, ttlgram, ttlbiaya int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	sisakg = gram % 1000
	sisaGram = sisakg
	ttlkg = kg * 10000

	if kg > 10 {
		sisaGram = 0
	} else {
		if sisaGram >= 500 {
			ttlgram = sisaGram * 5
		} else {
			ttlgram = sisaGram * 15
		}
	}
	ttlbiaya = ttlkg + ttlgram

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Println("Detail berat : ", kg, " kg + ", sisakg, " gram")
	fmt.Println("Detail biaya  : Rp.", ttlkg, " + Rp.", ttlgram)
	fmt.Println("Total biaya: Rp", ttlbiaya)
}
