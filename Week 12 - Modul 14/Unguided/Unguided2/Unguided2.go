package main

import "fmt"

const MAX = 1000

type arrInt [MAX]int

func selectionSortAsc(A *arrInt, n int) {
	var i, j, min, temp int

	for i = 0; i < n-1; i++ {
		min = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[min] {
				min = j
			}
		}

		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

func selectionSortDesc(A *arrInt, n int) {
	var i, j, max, temp int

	for i = 0; i < n-1; i++ {
		max = i

		for j = i + 1; j < n; j++ {
			if A[j] > A[max] {
				max = j
			}
		}

		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}
}

func main() {
	var daerah, m int
	var i, j, bil int
	var ganjil, genap arrInt
	var nGanjil, nGenap int

	fmt.Scan(&daerah)

	for i = 0; i < daerah; i++ {

		fmt.Scan(&m)

		nGanjil = 0
		nGenap = 0

		for j = 0; j < m; j++ {
			fmt.Scan(&bil)

			if bil%2 == 1 {
				ganjil[nGanjil] = bil
				nGanjil++
			} else {
				genap[nGenap] = bil
				nGenap++
			}
		}

		selectionSortAsc(&ganjil, nGanjil)
		selectionSortDesc(&genap, nGenap)

		for j = 0; j < nGanjil; j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j = 0; j < nGenap; j++ {
			fmt.Print(genap[j], " ")
		}

		fmt.Println()
	}
}
