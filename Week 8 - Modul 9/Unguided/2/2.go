package main

import (
	"fmt"
	"math"
)

type tabInt [100]int

func rataRata(arr tabInt, n int) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += float64(arr[i])
	}
	return sum / float64(n)
}

func standarDeviasi(arr tabInt, n int) float64 {
	mean := rataRata(arr, n)
	var sumSquares float64
	for i := 0; i < n; i++ {
		diff := float64(arr[i]) - mean
		sumSquares += diff * diff
	}
	return math.Sqrt(sumSquares / float64(n))
}

func main() {
	var arr tabInt
	var n int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan %d bilangan:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("\na. Keseluruhan isi array : ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("b. Indeks ganjil         : ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("c. Indeks genap          : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan kelipatan indeks (x): ")
	fmt.Scan(&x)
	fmt.Printf("   Indeks kelipatan %d      : ", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idxHapus int
	fmt.Print("e. Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&idxHapus)

	for i := idxHapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Print("   Array setelah dihapus : ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	mean := rataRata(arr, n)
	fmt.Printf("f. Rata-rata             : %.2f\n", mean)

	sd := standarDeviasi(arr, n)
	fmt.Printf("g. Standar Deviasi       : %.2f\n", sd)

	var cari, frek int
	fmt.Print("h. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frek++
		}
	}
	fmt.Printf("   Frekuensi bilangan %d muncul sebanyak %d kali\n", cari, frek)
}