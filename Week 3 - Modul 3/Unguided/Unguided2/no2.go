package main

import "fmt"

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func main() {
	var a, b, c int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&c)
	fmt.Println()
	fmt.Println("f(g(h(", a, "))) :", f(g(h(a))))
	fmt.Println("g(h(f(", b, "))) :", g(h(f(b))))
	fmt.Println("h(f(g(", c, "))) :", h(f(g(c))))
}
