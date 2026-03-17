package main

import "fmt"

func motor(x,y int) int{
	var tarif int

	if x>=0 && y<=17 {
		tarif=4000*(y-x)
	} else if x>17 && y<=24 {
		tarif=5000*(y-x)
	} else if x<17 && y>17 {
		tarif=4000*(17-x)+5000*(y-17)
	}

	return tarif
}

func mobil(x,y int) int{
	var tarif int

	if x>=0 && y<=17 {
		tarif=6000*(y-x)
	} else if x>17 && y<=24 {
		tarif=7000*(y-x)
	} else if x<17 && y>17 {
		tarif=6000*(17-x)+7000*(y-17)
	}

	return tarif
}

func main(){
	var kendaraan string
	var x,y,hasil,nomor,ttl int

	nomor=1
	for {
		fmt.Printf("*Kendaraan %d\n", nomor)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&kendaraan)
		
		if kendaraan=="-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&x)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&y)

		if kendaraan=="motor" {
			hasil=motor(x,y)
			fmt.Printf("Biaya parkir motor %d: %d\n", nomor, hasil)
		} else if kendaraan=="mobil" {
			hasil=mobil(x,y)
			fmt.Printf("Biaya parkir mobil %d: %d\n", nomor, hasil)
		} else {
			fmt.Println("Input tidak valid!")
		}
		fmt.Println("==========================================")
		fmt.Println()

		ttl=ttl+hasil
		nomor++
	}

	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", ttl)
	
}