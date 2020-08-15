package main

import (
	"fmt"
	"os"
)

var str string = "SEE"

var length int = len(str)

func main() {

	pos := 0
	var arr = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			ch := arr[i][j]
			res := ""
			if ch == str[pos] {
				res += string(ch)
				// fmt.Println(i, j)
				find(arr, pos+1, []int{i, j}, res)
			}
		}
	}
	fmt.Println(false)

}

//find next in board

func find(arr [][]byte, pos int, s []int, res string) {
	i, j := s[0], s[1]
	// fmt.Println(res, i, j)

	arr[i][j] = '.'

	// check up
	if i > 0 {
		// fmt.Println("u")
		ch := arr[i-1][j]
		if ch == str[pos] {
			res += string(ch)
			if res == str {
				fmt.Println(res, true)
				os.Exit(1)
			}
			find(arr, pos+1, []int{i - 1, j}, res)
		}
	}
	if i <= 1 { //check down
		// fmt.Println("d")

		ch := arr[i+1][j]
		// fmt.Println(ch)
		if ch == str[pos] {
			res += string(ch)
			if res == str {
				fmt.Println(res, true)
				os.Exit(1)
			}
			find(arr, pos+1, []int{i + 1, j}, res)
		}
	}
	if j > 0 { //check left
		// fmt.Println("l")

		ch := arr[i][j-1]
		if ch == str[pos] {
			res += string(ch)
			if res == str {
				fmt.Println(res, true)
				os.Exit(1)
			}
			find(arr, pos+1, []int{i, j - 1}, res)
		}
	}
	if j <= 2 { //check right
		// fmt.Println("r")

		ch := arr[i][j+1]
		if ch == str[pos] {
			res += string(ch)
			if res == str {
				fmt.Println(res, true)
				os.Exit(1)
			}
			find(arr, pos+1, []int{i, j + 1}, res)
		}
	}

	return

}
