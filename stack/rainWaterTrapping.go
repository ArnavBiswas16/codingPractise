package main

import (
	"fmt"

	"../../github.com/golang-collections/collections/stack"
)

func main() {

	arr := []int{5, 2, 0, 6, 3, 4, 0, 6}
	ngrArr := ngr(arr)
	nglArr := ngl(arr)
	fmt.Println(ngrArr)
	fmt.Println(nglArr)

	total := 0
	for i := range arr {
		m := min(nglArr[i], ngrArr[i], arr)
		if m > -1 {
			fmt.Println(m, m-arr[i])

			total += m - arr[i]

		}
	}
	fmt.Println(total)
}

func ngl(arr []int) []int {
	var ans []int

	s := stack.New()
	for i := range arr {
		if s.Len() == 0 {
			ans = append(ans, -1)
		} else {
			val, ok := s.Peek().(int)
			if ok && arr[val] > arr[i] {
				ans = append(ans, val)
			}
			if arr[val] <= arr[i] {
				for s.Len() > 0 && arr[val] <= arr[i] {
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
		s.Push(i)
	}
	return ans
}

func ngr(arr []int) []int {

	var ans []int
	s := stack.New()
	for i := len(arr) - 1; i >= 0; i-- {
		if s.Len() == 0 {
			ans = append(ans, -1)
		} else if s.Len() > 0 {
			val, ok := s.Peek().(int)
			if ok && arr[val] > arr[i] {
				ans = append(ans, val)
			}
			if arr[val] <= arr[i] {
				for s.Len() > 0 && arr[val] <= arr[i] {
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
		s.Push(i)
	}
	l := len(arr)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func min(a int, b int, arr []int) int {
	if a == -1 || b == -1 {
		return -1
	} else if arr[a] < arr[b] {
		return arr[a]
	} else {
		return arr[b]
	}
}
