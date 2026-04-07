package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	var tanggungan int
	switch tujuan {
	case "domestik":
		if jumlahHari > 3 {
			tanggungan = 3
		} else {
			tanggungan = jumlahHari
		}
	case "mancanegara":
		if jumlahHari > 8 {
			tanggungan = 8
		} else {
			tanggungan = jumlahHari
		}
	}
	return tanggungan
}

func biayaPerHari(jumlahMhs int) int {
	var biaya int
	biaya = (35000 * 2) + 250000 + 300000
	return biaya * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	tanggungan := tanggunganHari(lamaPerjalanan, tujuan)
	biaya := biayaPerHari(jumlahMhs)
	switch tujuan {
	case "domestik":
		*totalBiaya = float64(tanggungan * biaya)
	case "mancanegara":
		*totalBiaya = float64(tanggungan*biaya) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	perhitunganBiaya(jumlah, lama, tujuan, &biaya)
	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.2f\n", biaya)
}

// lama perjalana maksimal 3 hari untuk domestik dan 8 hari untuk mancanegara
// biaya yang harus dialokasikan per mahasiswa per hari untuk domestik untuk makan siang dan makan malam, biaya 1 kali makan adalah Rp. 35.000
// biaya penginapan untuk domestik adalah 250.000 (sudah termasuk makan pagi)
// uang saku sebesar 300.000 per mahasiswa untuk domestik
// biaya yang harus dialokasikan per mahasiswa per hari untuk mancanegara adalah 1,5 kali biaya per hari untuk domestik
