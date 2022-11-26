package main

import (
	"fmt"
	"sort"
)

func main() {
	angka := []int{1, 1, 1, 2, 3, 2, 4, 2}
	sort.Ints(angka)
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
	fmt.Println(jumlah)
}
