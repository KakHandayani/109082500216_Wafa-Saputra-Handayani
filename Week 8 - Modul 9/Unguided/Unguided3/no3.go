package main

import "fmt"

func main() {
	var klub1, klub2 string
	fmt.Print("Klub 1: ")
	fmt.Scan(&klub1)
	fmt.Print("Klub 2: ")
	fmt.Scan(&klub2)
	var skor1, skor2 int
	var hasil []string
	i := 1
	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skor1, &skor2)
		if skor1 < 0 || skor2 < 0 {
			break
		}
		if skor1 > skor2 {
			hasil = append(hasil, klub1)
		} else if skor2 > skor1 {
			hasil = append(hasil, klub2)
		} else {
			hasil = append(hasil, "Draw")
		}
		i++
	}
	for j := 0; j < len(hasil); j++ {
		fmt.Printf("Hasil %d : %s\n", j+1, hasil[j])
	}
	fmt.Println("Pertandingan selesai")
}
