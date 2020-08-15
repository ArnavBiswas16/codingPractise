package main

import "fmt"

var arr = []int{1, 2, 2, 4}
var n = len(arr)

type r []int

var res []r

func main() {

	var s []int
	sol(0, 0, s)
	fmt.Println(res)

}

func sol(pos int, len int, subset []int) {
	// fmt.Println(pos, len, subset)

	if pos == n {
		// fmt.Println("p")
		res = append(res, subset)
		// fmt.Println(subset)

		return
	}

	subset1 := append(subset, arr[pos])
	// fmt.Println(subset)
	sol(pos+1, len+1, subset1)
	sol(pos+1, len, subset)

}
