package main

import (
	"fmt"

	"../../github.com/golang-collections/collections/stack"
)

func main() {

	arr := []int{1, 2, 4, 3, 7}
	var ans []int

	s := stack.New()

	for _, v := range arr {
		if s.Len() == 0 {
			ans = append(ans, -1)
		} else {
			top, ok := s.Peek().(int)
			if ok && v > top {
				ans = append(ans, top)
			}
			if top > v {
				for s.Len() > 0 && top > v {
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
		s.Push(v)
	}
	fmt.Println(ans)
}
