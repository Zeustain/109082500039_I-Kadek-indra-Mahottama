package main

import "fmt"

const pi float64 = 3.14

func volume(r,t float64) float64{
	var volume float64

	volume=(pi*(r*r))*t

	return volume
}

func massa(r,t,p float64)float64{
	var massa float64

	massa=volume(r,t)*p

	return massa
}

func display(m1,m2 float64){
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		selisih := m1 - m2
		if selisih < 0 {
			selisih = selisih * (-1)
		} 
		fmt.Print("Selisih : ", selisih)
	}
}

func main(){
	var jarijari float64
	var t1, p1 float64
	var t2, p2 float64

	fmt.Print("Masukan jari-jari alas tabung : ")
	fmt.Scan(&jarijari)
	fmt.Println()
	fmt.Print("Masukan tinggi zat cair tabung kiri : ")
	fmt.Scan(&t1)
	fmt.Print("Masukan masa jenis zat cair tabung kiri : ")
	fmt.Scan(&p1)
	fmt.Println()
	fmt.Print("Masukan tinggi zat cair tabung kanan : ")
	fmt.Scan(&t2)
	fmt.Print("Masukan masa jenis zat cair tabung kanan : ")
	fmt.Scan(&p2)
	fmt.Println()

	m1 := massa(jarijari, t1, p1)
	m2 := massa(jarijari, t2, p2)

	display(m1, m2)
}