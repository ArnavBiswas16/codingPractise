package main

import (
	"fmt"

	"../../github.com/golang-collections/collections/stack"
)

func main() {

	arr := []int{9, 2, 4, 3, 7}
	var ans []int

	s := stack.New()

	for i := len(arr) - 1; i >= 0; i-- {
		if s.Len() == 0 {
			ans = append(ans, -1)
		} else {
			top, ok := s.Peek().(int)
			if ok && top < arr[i] {
				ans = append(ans, top)
			} else if top > arr[i] {
				for s.Len() > 0 && top > arr[i] {
					s.Pop()
					top, ok = s.Peek().(int)
				}

				if s.Len() == 0 {
					ans = append(ans, -1)
				} else {
					top, ok := s.Peek().(int)
					if ok {
						ans = append(ans, top)
					}
				}
			}
		}
		s.Push(arr[i])
	}
	fmt.Println(ans)
}
