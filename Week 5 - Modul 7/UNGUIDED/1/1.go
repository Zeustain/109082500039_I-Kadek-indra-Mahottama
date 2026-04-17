package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu{
	var hasil suhu

	hasil=(4*Celcius)/5

	return suhu(hasil)
}

func CelciusToFahrenheit(Celcius suhu) suhu{
	var hasil suhu

	hasil=(Celcius*9/5)+32

	return suhu(hasil)
}

func CelciusToKelvin(Celcius suhu) suhu{
	var hasil suhu

	hasil=Celcius+273.15

	return suhu(hasil)
}

func main(){
	var Celcius suhu
	var hasilR,hasilF,hasilK suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Maukan Suhu (Celcius) : ")
	fmt.Scan(&Celcius)

	hasilR=CelciusToReamur(Celcius)
	hasilF=CelciusToFahrenheit(Celcius)
	hasilK=CelciusToKelvin(Celcius)

	fmt.Println()
	fmt.Printf("%.f celcius = %v reamur\n", Celcius,hasilR)
	fmt.Printf("%.f celcius = %v fahrenheit\n", Celcius,hasilF)
	fmt.Printf("%.f celcius = %v kelvin", Celcius,hasilK)

}