package main

import "fmt"

type arrKelinci [1000]float64

func main() {
	var n int
	var berat arrKelinci
	var min, max float64

	fmt.Print("Jumlah kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("berat kelinci ke-",i+1," :")
		fmt.Scan(&berat[i])
	}

	if n > 0 {
		min = berat[0]
		max = berat[0]

		for i := 1; i < n; i++ {
			if berat[i] < min {
				min = berat[i]
			}
			if berat[i] > max {
				max = berat[i]
			}
		}

		fmt.Printf("berat terkecil: %v, berat terbesar: %v", min, max)
	}
}