package main

import (
	"fmt"

	"../../github.com/golang-collections/collections/stack"
)

func main() {

	arr := []int{100, 80, 60, 70, 60, 75, 85}
	var ans []int

	s := stack.New()
	for i, val := range arr {
		if s.Len() == 0 {
			ans = append(ans, -1)
		} else {
			v, ok := s.Peek().(int)
			if ok && arr[v] > val {
				ans = append(ans, v)
			} else if ok && arr[v] < val {
				for s.Len() > 0 && arr[v] < val {
					s.Pop()
					v, ok = s.Peek().(int)
				}
				if s.Len() == 0 {
					ans = append(ans, -1)
				} else {
					v, ok := s.Peek().(int)
					if ok {
						ans = append(ans, v)
					}
				}
			}
		}
		s.Push(i)
	}
	fmt.Println(ans)
	for i, val := range ans {
		if val == -1 {
			ans[i] = 1
		} else {
			ans[i] = i - val
		}
	}
	fmt.Println(ans)
}
