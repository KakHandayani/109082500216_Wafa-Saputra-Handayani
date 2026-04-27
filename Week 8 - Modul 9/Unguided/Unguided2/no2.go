package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println("\nIsi array:", arr)
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Array setelah dihapus:", arr)
	var total float64
	for _, v := range arr {
		total += float64(v)
	}
	rata := total / float64(len(arr))
	fmt.Println("Rata-rata:", rata)
	var jumlah float64
	for _, v := range arr {
		jumlah += math.Pow(float64(v)-rata, 2)
	}
	stdDev := math.Sqrt(jumlah / float64(len(arr)))
	fmt.Println("Standar deviasi:", stdDev)
	var cari, count int
	fmt.Print("Masukkan bilangan yang dicari: ")
	fmt.Scan(&cari)
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}
