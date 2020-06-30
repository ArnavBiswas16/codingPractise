package main

import "fmt"

func main() {

	var a = []int{19, 3, 7, 1, 2}
	fmt.Println(a)

	sortedArray := selectionSort(a)
	fmt.Println(sortedArray)

}

func selectionSort(a []int) []int {

	for i := 0; i < len(a); i++ {
		leastInd := i

		var j int
		for j = i + 1; j < len(a); j++ {

			if a[j] < a[leastInd] {
				leastInd = j
			}
		}
		temp := a[i]
		a[i] = a[leastInd]
		a[leastInd] = temp
	}
	return a
}
