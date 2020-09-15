package main

import (
	"fmt"
	"sort"
)

func main() {

	intervals := [][]int{{12, 13}, {2, 8}, {1, 3}, {6, 11}}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	i := 0

	for i < len(intervals) {
		left, right := intervals[i][0], intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= right {
			right = max(right, intervals[j][1])
			j++
		}
		res = append(res, []int{left, right})
		i = j
	}
	fmt.Println(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
