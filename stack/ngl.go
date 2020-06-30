package main

import (
	"fmt"

	"../../github.com/golang-collections/collections/stack"
)

func main() {

	arr := []int{3, 2, 4, 1}
	var ans []int

	s := stack.New()
	for i := range arr {
		if s.Len() == 0 {
			ans = append(ans, -1)
		} else {
			val, ok := s.Peek().(int)
			if ok && val > arr[i] {
				ans = append(ans, val)
			}
			if val <= arr[i] {
				for s.Len() > 0 && val < arr[i] {
					s.Pop()
					val, ok = s.Peek().(int)
				}

				if s.Len() == 0 {
					ans = append(ans, -1)
				} else {
					val, ok := s.Peek().(int)
					if ok {
						ans = append(ans, val)
					}
				}
			}
		}
		s.Push(arr[i])
	}

	fmt.Println(ans)
}
