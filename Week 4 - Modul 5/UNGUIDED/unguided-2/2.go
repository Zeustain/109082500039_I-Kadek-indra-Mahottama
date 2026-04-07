package main

import "fmt"

func angka(n int){
	if n==1 {
		fmt.Print(n," ")
		return 
	}
	fmt.Print(n," ")
	angka(n-1)
	fmt.Print(n," ")
}

func main(){
	var n int

	fmt.Print("Masukan angka: ")
	fmt.Scan(&n)

	angka(n)
}
