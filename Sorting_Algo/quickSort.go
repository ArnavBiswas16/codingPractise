package main

import "fmt"

func main() {

	var a = []int{10, 16, 8, 12, 15, 6, 3, 9, 5}
	fmt.Println(a)
	a = append(a, 999999999)

	sortedArray := quickSort(a, 0, len(a)-1)
	sortedArray = sortedArray[:len(a)-1]
	fmt.Println(sortedArray)

}

func quickSort(a []int, lb int, ub int) []int {
	if lb < ub {
		loc := partition(a, lb, ub)
		quickSort(a, lb, loc)
		quickSort(a, loc+1, ub)
	}

	return a
}

func partition(a []int, lb int, ub int) int {

	pivot := a[lb]
	start := lb
	end := ub

	for start < end {

		for {
			start++
			if a[start] > pivot {
				break
			}
		}
		for {
			end--
			if a[end] <= pivot {
				break
			}

		}
		if start < end {
			temp := a[end]
			a[end] = a[start]
			a[start] = temp
		}

	}
	temp := a[end]
	a[end] = a[lb]
	a[lb] = temp
	return end
}
