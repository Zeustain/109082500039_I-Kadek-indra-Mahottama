package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var c rune
	*n = 0
	for {
		fmt.Scanf("%c", &c)
		if c == '.' || *n >= NMAX {
			break
		}
		if c != ' ' && c != '\n' && c != '\r' {
			t[*n] = c
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	tReverse := t
	balikanArray(&tReverse, n)
	for i := 0; i < n; i++ {
		if t[i] != tReverse[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Masukkan karakter teks (akhiri dengan titik '.'): \n")
	isiArray(&tab, &m)

	fmt.Print("Teks asli    : ")
	cetakArray(tab, m)

	fmt.Print("Reverse teks : ")
	balikanArray(&tab, m)
	cetakArray(tab, m)
	
	balikanArray(&tab, m)

	hasilPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom    : %t\n", hasilPalindrom)
}