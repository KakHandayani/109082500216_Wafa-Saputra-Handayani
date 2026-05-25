package main

import "fmt"

const MAX = 1000

type arrInt [MAX]int

func insertionSort(A *arrInt, n int) {
	var i, pass, temp int

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass

		for i > 0 && temp < A[i-1] {
			A[i] = A[i-1]
			i--
		}

		A[i] = temp
	}
}

func jarakTetap(A arrInt, n int) bool {
	var selisih int
	var i int

	if n <= 2 {
		return true
	}

	selisih = A[1] - A[0]

	for i = 2; i < n; i++ {
		if A[i]-A[i-1] != selisih {
			return false
		}
	}

	return true
}

func main() {
	var A arrInt
	var n int
	var x int
	var i int

	n = 0

	fmt.Scan(&x)

	for x >= 0 {
		A[n] = x
		n++

		fmt.Scan(&x)
	}

	insertionSort(&A, n)

	for i = 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	if jarakTetap(A, n) {
		fmt.Println("Data berjarak tetap")
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
