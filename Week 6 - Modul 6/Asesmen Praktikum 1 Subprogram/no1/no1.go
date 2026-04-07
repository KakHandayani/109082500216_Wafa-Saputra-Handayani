package main

import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64 {
	return (1.0 / 3.0) * pi * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	fmt.Printf("Massa jenis air: %.2f kg/m^3\n", m1)
	fmt.Printf("Massa jenis minyak: %.2f kg/m^3\n", m2)
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var mjKiri, mjKanan float64
	var massaKiri, massaKanan float64
	fmt.Print("Masukkan jari-jari alas tabung: ")
	fmt.Scan(&r)

	fmt.Print("\nMasukkan tinggi zat cair tabung kiri: ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri: ")
	fmt.Scan(&mjKiri)

	fmt.Print("\nMasukkan tinggi zat cair tabung kanan: ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan: ")
	fmt.Scan(&mjKanan)
	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)
	if massaKiri == massaKanan {
		fmt.Println("\nBALANCE")
	} else {
		fmt.Printf("\nselisih massa zat cair kiri dan massa zat cair kanan: %.2f kg", massaKiri-massaKanan)
	}
}
