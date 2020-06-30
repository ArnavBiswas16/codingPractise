package main

import "fmt"

func main() {

	var a = []int{19, 3, 7, 1, 2}
	fmt.Println(a)

	sortedArray := insertionSort(a)
	fmt.Println(sortedArray)

}

func insertionSort(a []int) []int {

	for i := 1; i < len(a); i++ {
		temp := a[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if a[j] > temp {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = temp
	}
	return a
}
