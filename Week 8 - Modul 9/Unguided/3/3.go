package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang [100]string
	var n int

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	n = 0
	for {
		fmt.Printf("Pertandingan %d: ", n+1)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[n] = klubA
		} else if skorB > skorA {
			pemenang[n] = klubB
		} else {
			pemenang[n] = "Draw"
		}
		
		n++
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	
	fmt.Println("Pertandingan selesai")
}