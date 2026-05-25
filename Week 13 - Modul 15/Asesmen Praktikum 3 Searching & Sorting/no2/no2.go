package main

import "fmt"

const NMAX = 1005

type Pemain struct {
	NamaDepan    string
	NamaBelakang string
	Gol          int
	Assist       int
}
type DaftarPemain [NMAX]Pemain

func InsertionSort(T *DaftarPemain, n int) {
	for i := 1; i < n; i++ {
		key := T[i]
		j := i - 1
		for j >= 0 && (T[j].Gol < key.Gol || (T[j].Gol == key.Gol && T[j].Assist < key.Assist)) {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}
func main() {
	var n int
	var data DaftarPemain
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].NamaDepan, &data[i].NamaBelakang, &data[i].Gol, &data[i].Assist)
	}
	InsertionSort(&data, n)
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", data[i].NamaDepan, data[i].NamaBelakang, data[i].Gol, data[i].Assist)
	}
}
