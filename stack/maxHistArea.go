package main

import (
	"fmt"

	"../../github.com/golang-collections/collections/stack"
)

func main() {

	arr := []int{0, 1, 1, 0}
	nsr := nsr(arr)
	nsl := nsl(arr)

	maxArea := 0
	for i := 0; i < len(arr); i++ {
		val := (nsr[i] - nsl[i]) * arr[i]
		if val > maxArea {
			maxArea = val
		}
	}
	fmt.Println(maxArea)
}

func nsl(arr []int) []int {
	var nsl []int
	s := stack.New()
	for i, val := range arr {
		if s.Len() == 0 {
			nsl = append(nsl, 0)
		} else {
			top, ok := s.Peek().(int)

			if ok && arr[top] < val {
				nsl = append(nsl, top+1)
			} else if arr[top] >= val {
				for s.Len() > 0 && arr[top] >= val {
					s.Pop()
					top, ok = s.Peek().(int)
				}

				if s.Len() == 0 {
					nsl = append(nsl, 0)
				} else {
					top, ok := s.Peek().(int)
					if ok {
						nsl = append(nsl, top+1)
					}
				}
			}
		}
		s.Push(i)
	}
	return nsl
}

func nsr(arr []int) []int {
	s := stack.New()
	var ans []int
	l := len(arr)
	for i := l - 1; i >= 0; i-- {
		if s.Len() == 0 {
			ans = append(ans, l)
		} else {
			top, ok := s.Peek().(int)
			if ok && arr[top] < arr[i] {
				ans = append(ans, top)
			} else if arr[top] >= arr[i] {
				for s.Len() > 0 && arr[top] >= arr[i] {
					s.Pop()
					top, ok = s.Peek().(int)
				}

				if s.Len() == 0 {
					ans = append(ans, l)
				} else {
					top, ok := s.Peek().(int)
					if ok {
						ans = append(ans, top)
					}
				}
			}
		}
		s.Push(i)
	}
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
