package main

import (
	"fmt"
	"../../github.com/golang-collections/collections/stack"
)

func main() {

	arr := []int{1, 3, 2, 4}
	var ans []int
	s := New()
	for i := len(arr) - 1; i >= 0; i-- {
		if s.Len() == 0 {
			ans = append(ans, -1)
			s.Push(arr[i])
		} else {
			for s.Peek < a[i] {
				s.Pop()
			}
			if s.Len() == 0 {
				ans = append(ans, -1)
				s.Push(arr[i])
			} else {
				ans = append(ans, s.Peek)
			}
		}
	}
	fmt.Println(ans)
}
