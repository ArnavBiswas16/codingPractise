// https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/
package main

import "fmt"

func main() {

	arr := []int{0, 0, 2, 1, 2, 1, 0, 2, 1}

	var c0, c1, c2, i int

	for _, val := range arr {
		switch val {
		case 0:
			c0++
			break
		case 1:
			c1++
			break
		case 2:
			c2++
			break
		}
	}

	for c0 > 0 {
		arr[i] = 0
		i++
		c0--
	}
	for c1 > 0 {
		arr[i] = 1
		i++
		c1--
	}
	for c2 > 0 {
		arr[i] = 2
		i++
		c2--
	}

	fmt.Println(arr)
}
