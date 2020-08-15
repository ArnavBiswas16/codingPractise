package main

import "fmt"

type Stack []int

func main() {

	var stack Stack
	area := 0
	maxArea := 0
	top := -1
	hist := [5]int{1, 2, 3, 4, 5}
	var i, h int
	for i, h = range hist {

		if (stack.IsEmpty()) || (h >= hist[stack[top]]) {
			stack.Push(i)
			top++

		} else {

			for h <= hist[stack[top]] {
				p, _ := stack.Pop()
				if stack.IsEmpty() {
					area = hist[p] * i
					if area > maxArea {
						maxArea = area
					}
					top--

					break
				} else {
					top--
					area = hist[p] * (i - stack[top] - 1)
					if area > maxArea {
						maxArea = area
					}
				}
			}
			stack.Push(i)
			top++

		}
		if (i == len(hist)-1) && (stack.IsEmpty() == false) {
			i++
			for stack.IsEmpty() == false {
				p, _ := stack.Pop()
				if stack.IsEmpty() {
					area = hist[p] * i
					if area > maxArea {
						maxArea = area
					}
					top--
					break
				} else {
					top--
					area = hist[p] * (i - stack[top] - 1)
					if area > maxArea {
						maxArea = area
					}
				}
			}
		}
	}

	fmt.Println(maxArea)
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(n int) {
	*s = append(*s, n) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		// fmt.Println(index)
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		// fmt.Println(element)
		return element, true
	}
}
