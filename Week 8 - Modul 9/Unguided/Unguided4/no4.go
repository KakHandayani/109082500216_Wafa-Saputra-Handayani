package main

import "fmt"

const MAX int = 127

type tabel [MAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= MAX {
			break
		}
		if ch != ' ' && ch != '\n' {
			t[*n] = ch
			*n++
		}
	}
}
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}
func main() {
	var tab tabel
	var n int
	fmt.Print("Hves : ")
	isiArray(&tab, &n)
	fmt.Print("Rvvsus : ")
	cetakArray(tab, n)
	if palindrom(tab, n) {
		fmt.Println("flalitaouo : true")
	} else {
		fmt.Println("flalitaouo : false")
	}
	balikanArray(&tab, n)
	fmt.Print("Hves : ")
	cetakArray(tab, n)
}
