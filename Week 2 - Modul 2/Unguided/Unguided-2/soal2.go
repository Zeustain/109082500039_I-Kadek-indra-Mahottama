package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	var wb1, wb2, wb3, wb4 string

	wb1 = "merah"
	wb2 = "kuning"
	wb2 = "hijau"
	wb2 = "ungu"

	Percobaan := 0
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %.d : ", i)
		fmt.Scan(&w1, &w2, &w3, &w4)
		if w1 != wb1 || w2 != wb2 || w3 != wb3 || w4 != wb4 {
			Percobaan++
		}
	}
	if Percobaan != 0 {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
