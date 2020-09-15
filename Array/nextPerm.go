//https://www.youtube.com/watch?v=zGQq3HGBTXg
package main

import "fmt"

func main() {

	arr := []int{9, 1, 2, 4, 3, 1, 0}
	fmt.Println(nextP(arr))
}

func nextP(arr []int) []int {

	len := len(arr)
	max := arr[len-1]
	inverseP := -1
	for i := len - 1; i >= 0; i-- {

		if arr[i] >= max {
			max = arr[i]
		} else {
			inverseP = i
			break
		}
	}

	if inverseP < 0 {
		return []int{}
	}

	for i := inverseP + 1; i < len; i++ {

		if arr[inverseP] > arr[i] {
			arr[inverseP], arr[i-1] = arr[i-1], arr[inverseP]
			break
		}
	}
	return append(arr[:inverseP+1], reverse(arr[inverseP+1:])...)
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
