package main

import "fmt"

type arrIkan [1000]float64

func main() {
	var x, y int
	var ikan arrIkan

	fmt.Print("Masukan jumlah ikan dijual dan banyak ikan didalam wadah: ")
	fmt.Scan(&x, &y)

	fmt.Print("berat ikan yang akan dijual: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	if x > 0 && y > 0 {
		var totalSemua float64
		var jumlahWadah int

		for i := 0; i < x; i += y {
			var beratWadah float64
			
			for j := 0; j < y && i+j < x; j++ {
				beratWadah += ikan[i+j]
			}
			
			fmt.Printf("%.2f  ", beratWadah)
			
			totalSemua += beratWadah
			jumlahWadah++
		}
		fmt.Println()

		rataRata := totalSemua / float64(jumlahWadah)
		fmt.Printf("berat rata rata : %.2f \n", rataRata)
	}
}