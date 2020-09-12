//https://www.youtube.com/watch?v=YxuK6A3SvTs&t=210s
package main

import (
	"fmt"
	"math/bits"
)

func main() {

	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	const (
		MaxInt = 1<<(bits.UintSize-1) - 1
		MinInt = -MaxInt - 1
	)
	meh := 0
	msf := MinInt

	for _, val := range arr {
		meh += val
		if meh < val {
			meh = val
		}
		if msf < meh {
			msf = meh
		}
	}
	fmt.Println(msf)
}
