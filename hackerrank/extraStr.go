package main

import "fmt"

func main() {
	a := "abcc"
	b := "abccc"

	var x rune
	for _, c := range a {
		x = x + c
	}
	var y rune
	for _, c := range b {
		y = y + c
	}
	result := y - x
	resCh := string(result)
	fmt.Println(resCh)
}
