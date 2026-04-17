package main

import "fmt"

type angka int
type kata string
type Buku struct{
	judul kata
	penulis kata
	penerbit kata
	tahunTerbit angka
	jumlahHalaman angka
}

func main(){
	var bio Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukan judul buku : ")
	fmt.Scan(&bio.judul)
	fmt.Print("Masukan nama penulis : ")
	fmt.Scan(&bio.penulis)
	fmt.Print("Masukan penerbit : ")
	fmt.Scan(&bio.penerbit)
	fmt.Print("Masukan tahun terbit : ")
	fmt.Scan(&bio.tahunTerbit)
	fmt.Print("Masukan jumlah halaman : ")
	fmt.Scan(&bio.jumlahHalaman)

	fmt.Println()

	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("Judul Buku : ", bio.judul)
	fmt.Println("Penulis : ", bio.penulis)
	fmt.Println("Penerbit : ", bio.penerbit)
	fmt.Println("Tahun Terbit : ", bio.tahunTerbit)
	fmt.Println("Jumlah Halaman : ", bio.jumlahHalaman)
}