package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return c * 4 / 5
}
func CelciusToFahrenheit(c suhu) suhu {
	return (c * 9 / 5) + 32
}
func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}
func main() {
	var c suhu
	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&c)
	fmt.Println()
	fmt.Println(c, "celcius =", CelciusToReamur(c), "reamur")
	fmt.Println(c, "celcius =", CelciusToFahrenheit(c), "fahrenheit")
	fmt.Println(c, "celcius =", CelciusToKelvin(c), "kelvin")
}
