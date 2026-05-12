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
	ketua := 1
	wakil := 1
	for i := 2; i <= 20; i++ {
		if calon[i] > calon[ketua] {
			wakil = ketua
			ketua = i
		} else if calon[i] > calon[wakil] && i != ketua {
			wakil = i
		}
	}
	for i := 1; i <= 20; i++ {
		if i != ketua {
			if calon[i] > calon[wakil] {
				wakil = i
			}
		}
	}
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println()
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println()
	fmt.Println("Ketua RT:", ketua)
	fmt.Println()
	fmt.Println("Wakil ketua:", wakil)
}
