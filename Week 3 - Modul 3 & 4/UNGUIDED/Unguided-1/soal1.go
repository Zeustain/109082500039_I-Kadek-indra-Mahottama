package main

import "fmt"

func factorial(n int) int {
	var hasil int

	hasil = 1
	for i := n; i >= 1; i-- {
		hasil = hasil * i
	}

	return hasil
}

func permutation(n, r int) int {
	var hasil int

	hasil = factorial(n) / factorial(n-r)

	return hasil
}

func combination(n, r int) int {
	var hasil int

	hasil = factorial(n) / (factorial(r) * factorial(n-r))

	return hasil
}

func main() {
	var a, b, c, d int

	fmt.Println("Masukan empat angka: ")
	fmt.Scan(&a,&b,&c,&d)

	perm1 := permutation(a, c)
	com1 := combination(a, c)
	perm2 := permutation(b, d)
	com2 := combination(b, d)

	fmt.Println("==== Output ===")
	fmt.Println(perm1, com1)
	fmt.Println(perm2, com2)

}
