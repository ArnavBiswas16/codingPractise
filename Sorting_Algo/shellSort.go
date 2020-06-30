package main

import "fmt"

func main() {

	var a = []int{23, 29, 15, 19, 31, 7, 9, 5, 2}
	fmt.Println(a)

	sortedArray := shellSort(a)
	fmt.Println(sortedArray)

}

func shellSort(a []int) []int {
	gap := len(a) / 2

	for gap >= 1 {
		i := 0
		j := gap

		for j <= len(a)-1 {
			if a[i] > a[j] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
				reverseGap := i - gap
				ri := i
				for reverseGap >= 0 {
					if a[reverseGap] > a[ri] {
						temp := a[ri]
						a[ri] = a[reverseGap]
						a[reverseGap] = temp
					}
					ri--
					reverseGap--
				}

			}
			i++
			j++
		}
		gap = gap / 2
	}
	return a
}
