// https://www.geeksforgeeks.org/find-duplicates-in-on-time-and-constant-extra-space/
package main

import "fmt"

func main() {

	arr := [...]int{0, 4, 3, 2, 7, 8, 2, 3, 1}
	len := len(arr)
	for _, val := range arr {

		arr[val%len] += len
	}

	for i, val := range arr {
		if val >= len*2 {
			fmt.Print(i, ", ")
		}
	}

}
