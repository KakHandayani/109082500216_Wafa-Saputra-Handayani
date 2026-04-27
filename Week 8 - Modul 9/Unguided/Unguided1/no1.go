package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}
type lingkaran struct {
	pusat titik
	r     float64
}

func jarak(p, q titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}
func diDalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= c.r
}
func main() {
	var c1, c2 lingkaran
	var p titik
	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.r)
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.r)
	fmt.Scan(&p.x, &p.y)
	dalam1 := diDalam(c1, p)
	dalam2 := diDalam(c2, p)
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
