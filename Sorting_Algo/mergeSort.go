package main

import "fmt"

func main() {

	var a = []int{15, 5, 24, 8, 1, 3, 16, 10, 20}
	fmt.Println(a)

	sortedArray := mergeSort(a, 0, len(a)-1)
	fmt.Println(sortedArray)

}

func mergeSort(a []int, lb int, ub int) []int {
	if lb < ub {
		mid := (lb + ub) / 2
		mergeSort(a, lb, mid)
		mergeSort(a, mid+1, ub)
		a = merge(a, lb, mid, ub)
	}

	return a
}

func merge(a []int, lb int, mid int, ub int) []int {

	i := lb
	j := mid + 1
	k := lb
	b := make([]int, ub+1)
	for i <= mid && j <= ub {
		if a[i] <= a[j] {
			b[k] = a[i]
			i++
		} else {

			b[k] = a[j]
			j++
		}
		k++
	}

	if i > mid {
		for j <= ub {
			b[k] = a[j]
			k++
			j++
		}
	} else if j > ub {
		for i <= mid {
			b[k] = a[i]
			i++
			k++
		}
	}
	for k = lb; k <= ub; k++ {
		a[k] = b[k]
	}
	return a

}
