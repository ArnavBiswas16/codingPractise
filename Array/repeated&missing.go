//https://www.geeksforgeeks.org/find-a-repeating-and-a-missing-number/
package main

import "fmt"

func main() {

	arr := []int{7, 3, 4, 5, 5, 6, 2}

	for _, val := range arr {

		if val < 0 {
			val = -val
		}
		if arr[val-1] > 0 {
			arr[val-1] = -arr[val-1]
		} else {
			fmt.Println("repeated val =", val)
		}
	}

	for i, val := range arr {
		if val > 0 {
			fmt.Println("missing val = ", i+1)
			break
		}
	}

}
