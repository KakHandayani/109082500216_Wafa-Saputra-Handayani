package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   int
	nama  string
	nilai int
}
type arrayMahasiswa [nMax]mahasiswa

func inputData(arr *arrayMahasiswa, n *int) {
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data ke-%d: ", i+1)
		fmt.Scan(&arr[i].NIM, &arr[i].nama, &arr[i].nilai)
	}
}
func nilaiPertama(arr arrayMahasiswa, n int, nim int) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			return arr[i].nilai
		}
	}
	return -1
}
func nilaiTerbesar(arr arrayMahasiswa, n int, nim int) int {
	max := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			if arr[i].nilai > max {
				max = arr[i].nilai
			}
		}
	}
	return max
}
func main() {
	var data arrayMahasiswa
	var n int
	var nimCari int
	inputData(&data, &n)
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya: ")
	fmt.Scan(&nimCari)
	np := nilaiPertama(data, n, nimCari)
	nt := nilaiTerbesar(data, n, nimCari)
	fmt.Printf("Nilai pertama dari NIM %d adalah %d\n", nimCari, np)
	fmt.Printf("Nilai terbesar dari NIM %d adalah %d\n", nimCari, nt)
}
