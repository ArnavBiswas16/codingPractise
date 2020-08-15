package main

import (
	"fmt"
)

func main() {

	str := "cat is running"
	fmt.Println(str)
	str = "  " + str
	rev(str, "")
}

func rev(str string, res string) {

	if str == "" || len(str) <= 1 {
		res += str
		fmt.Print(res)
		return
	}

	c := string(str[len(str)-1])
	if c == " " {
		rev(res, " ")
		res = ""
	} else {
		res += c
	}
	rev(str[0:len(str)-1], res)

}
