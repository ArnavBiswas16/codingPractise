package main

import "fmt"

type Stack []string

func main() {
	s := "{[]}"

	result := isBalanced(s)

	fmt.Println(result)
}

// Complete the isBalanced function below.
func isBalanced(s string) string {

	if len(s)%2 != 0 {
		return "NO"
	}
	bracketMap := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	var stack Stack
	for _, c := range s {

		char := string(c)
		if char == "{" || char == "[" || char == "(" {
			stack.Push(char)
		} else if char == "}" || char == "]" || char == ")" {
			bracket, x := stack.Pop()
			if x == true {
				b := bracketMap[bracket]
				if char != b {
					return "NO"
				}
			} else {
				return ""
			}
		}
	}
	return "yes"
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
