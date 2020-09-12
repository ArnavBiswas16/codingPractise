package main

import (
	"fmt"
	"math"
)

func main() {

	arr1 := []int{3, 27, 38, 43}
	arr2 := []int{9, 10, 82}

	l1 := len(arr1)
	l2 := len(arr2)
	t := l1 + l2
	var GAP int = int(math.Ceil(float64(t) / 2))

	for GAP > 0 {
		var elem1, elem2 *int
		for p1 := 0; p1 < t-GAP; p1++ {
			// fmt.Println("p1", p1)
			if p1 >= l1 {
				elem1 = &arr2[p1-l1]
			} else {
				elem1 = &arr1[p1]
			}

			p2 := p1 + GAP
			// fmt.Println("p2", p2)

			if p2 >= l1 {
				elem2 = &arr2[p2-l1]
			} else if p2 < t {
				// fmt.Println(p2)
				elem2 = &arr1[p2]
			}
			// fmt.Println(*elem1, *elem2)

			if *elem1 > *elem2 {
				temp := *elem2
				*elem2 = *elem1
				*elem1 = temp
			}
		}

		GAP = GAP / 2

	}

	fmt.Println(arr1, arr2)
}
