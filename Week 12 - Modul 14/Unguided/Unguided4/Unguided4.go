package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	fmt.Println("Data buku terfavorit:")

	fmt.Println(
		pustaka[0].judul,
		pustaka[0].penulis,
		pustaka[0].penerbit,
		pustaka[0].tahun,
	)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var pass, i int
	var temp Buku

	for pass = 1; pass < n; pass++ {
		temp = pustaka[pass]
		i = pass

		for i > 0 && temp.rating > pustaka[i-1].rating {
			pustaka[i] = pustaka[i-1]
			i--
		}

		pustaka[i] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	fmt.Println("5 Buku dengan rating tertinggi:")

	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	for i = 0; i < batas; i++ {
		fmt.Println(
			pustaka[i].judul,
			"- Rating:",
			pustaka[i].rating,
		)
	}
}

func CariBuku(pustaka DaftarBuku, n, x int) {
	var kiri, kanan, tengah int
	var found bool

	kiri = 0
	kanan = n - 1
	found = false

	for kiri <= kanan && !found {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == x {
			found = true
		} else if x > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if found {
		fmt.Println("Data buku ditemukan:")
		fmt.Println("Judul :", pustaka[tengah].judul)
		fmt.Println("Penulis :", pustaka[tengah].penulis)
		fmt.Println("Penerbit :", pustaka[tengah].penerbit)
		fmt.Println("Tahun :", pustaka[tengah].tahun)
		fmt.Println("Rating :", pustaka[tengah].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, x int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan data buku:")
	DaftarkanBuku(&pustaka, n)

	UrutBuku(&pustaka, n)

	fmt.Println()
	CetakTerfavorit(pustaka, n)

	fmt.Println()
	Cetak5Terbaru(pustaka, n)

	fmt.Print("\nMasukkan rating buku yang dicari: ")
	fmt.Scan(&x)

	CariBuku(pustaka, n, x)
}
