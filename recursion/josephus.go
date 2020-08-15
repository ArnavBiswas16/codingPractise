package main

import "fmt"

func main() {

	n := 40
	k := 7
	var a []int

	for i := 1; i <= n; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	r := kill(a, k-1, 0)
	fmt.Println(r)
}

func kill(a []int, k int, pos int) []int {

	if len(a) == 1 {
		return a
	}

	pos += k
	pos = pos % (len(a))
	a = eliminate(a, pos)
	return kill(a, k, pos)
}

func eliminate(a []int, i int) []int {

	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = 0      // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a
}
