package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var kata = "Masing-masing anak mendap(atkan uang jajan ya=ng be&rbeda."
	var string = strings.Split(kata, " ")

	var regex, _ = regexp.Compile(`[a-z,A-Z,?,-.]+`)
	var jumlah = 0

	for i := 0; i < len(string); i++ {
		string[i] = regex.ReplaceAllString(string[i], "*")
		if string[i] == "*" {
			jumlah += 1
		}
	}
	fmt.Println(jumlah)
}
