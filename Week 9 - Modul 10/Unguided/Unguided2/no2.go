package main

import "fmt"

func main() {
	const MAX = 1000
	var arr [MAX]float64
	var x, y int
	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}
	var wadah [MAX]float64
	var count int = 0
	for i := 0; i < x; i += y {
		sum := 0.0
		for j := i; j < i+y && j < x; j++ {
			sum += arr[j]
		}
		wadah[count] = sum
		count++
	}
	for i := 0; i < count; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
	fmt.Println()
	total := 0.0
	for i := 0; i < count; i++ {
		total += wadah[i]
	}
	rata := total / float64(count)
	fmt.Printf("%.2f\n", rata)
}
