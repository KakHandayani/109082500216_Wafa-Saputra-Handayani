package main

import "fmt"

func main() {
	var suara int
	var totalMasuk, suaraSah int
	var calon [21]int
	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		totalMasuk++
		if suara >= 1 && suara <= 20 {
			suaraSah++
			calon[suara]++
		}
	}
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println()
	for i := 1; i <= 20; i++ {
		if calon[i] > 0 {
			fmt.Printf("%d: %d\n", i, calon[i])
		}
	}
}
