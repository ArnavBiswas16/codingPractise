package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter the denominations , separated :")
	var denom string
	fmt.Scanln(&denom)
	den := stringToIntArr(denom)

	fmt.Print("Enter the change :")
	var change int
	fmt.Scanln(&change)

	var l = len(den)
	// var a [][]int
	a := make([][]int, l)
	for p := range a {
		a[p] = make([]int, change+1)
	}

	for i := 0; i < l; i++ {

		for j := 0; j <= change; j++ {
			if j == 0 {
				a[i][j] = 1
				continue
			}
			if i == 0 {
				if j%den[i] == 0 {
					a[i][j] = 1
				} else {
					a[i][j] = 0
				}
			} else {
				if den[i] > j {
					a[i][j] = a[i-1][j]
				} else {

					a[i][j] = a[i-1][j] + a[i][j-den[i]]
				}
			}
		}
	}
	fmt.Println(a[l-1][change])
}

func stringToIntArr(t string) []int {

	var t2 = []int{}
	strArr := strings.Split(t, ",")

	for _, i := range strArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
}
