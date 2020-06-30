package main

import "fmt"

func main() {

	var a = []int{19, 3, 7, 1, 2}

	sortedArray := bubbleSort(a)
	fmt.Println(sortedArray)

}

func bubbleSort(a []int) []int {

	size := len(a)
	for i := 0; i < (size - 1); i++ {
		flag := 0
		for j := 0; j < (size - i - 1); j++ {

			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				flag = 1
			}
		}

		if flag == 0 {
			break
		}
	}
	return a
}
