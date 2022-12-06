package main

import (
	"fmt"
	"sort"
)

func main() {
	// angka := []int{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	angka := []int{}
	var nilai int

	for i := 0; i < 6; i++ {
		fmt.Scanln(&nilai)
		angka = append(angka, nilai)
	}
	sort.Ints(angka)
	// // [1,1,1,2,2,2,3,4]
	var jumlah = 0
	y := 1
	for i := 0; i < len(angka); i++ {
		if y != len(angka) {
			if angka[i] == angka[y] {
				jumlah += 1
				i += 1
				y += 2
			} else {
				y += 1
			}
		}
	}
	fmt.Println("Terdapat", jumlah, "Pasangan ")
}
