package main

import "fmt"

func main() {

	arr := []int{4, 2, 1, 3}
	count := 0
	i := 0
	for i < len(arr) {
		if arr[i] == i+1 {
			i++
		} else {
			temp := arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
			count++

		}
	}
	fmt.Println(count)
}
