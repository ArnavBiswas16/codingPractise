// A COVID center in Delhi inaugurated is by "Mota Bhai".
// This hall accommodates M*N beds. Currently all the beds are occupied by different people of different age groups.
// Today they are stuck due to COVID-19 lock down imposed by "NaMo".
// It is assumed that person in an age group of 18 to 50 including both is having good immunity.
// if one person gets infected (Suppose some Mufler Man with low immunity) in the hall then (s)he may infect the
//  person who are in his/her neighborhood and who are vulnerable. Find how many more persons will get infected if one person
//  from the hall will get infected.

// Age is between 1 and 100.
// Complete age are given in integer.
// All beds in M * N are occupied.

// Input:
// Matrix of age
// position of person

// Output:
// No of person who will get affected and are vulnerable. In short whose chances are high ki wo nahi bachenge" That they will not survive".

// Example
// 02 42 56 15
// 42 22 10 34
// 67 89 45 78
// 50 45 57 23
// 05 45 60 23

// Infected is person with age 34.

package main

import "fmt"

func main() {
	arr := [][]int{
		{2, 42, 56, 15},
		{42, 22, 10, 34},
		{67, 89, 45, 78},
		{50, 45, 57, 23},
		{5, 45, 60, 23},
	}

	s := []int{1, 3}

	find(arr, s)
	fmt.Println(c)
}

var c int

func find(arr [][]int, s []int) {

	i := s[0]
	j := s[1]

	arr[i][j] = 0

	// check up
	if i > 0 {
		val := arr[i-1][j]
		if val != 0 && (val < 18 || val > 50) {
			c++
			find(arr, []int{i - 1, j})
		}
	}
	if i <= 3 { //check down
		val := arr[i+1][j]
		if val != 0 && (val < 18 || val > 50) {
			c++
			find(arr, []int{i + 1, j})
		}
	}
	if j > 0 { //check left

		val := arr[i][j-1]
		if val != 0 && (val < 18 || val > 50) {
			c++
			find(arr, []int{i, j - 1})
		}
	}
	if j <= 2 { //check right
		val := arr[i][j+1]
		if val != 0 && (val < 18 || val > 50) {
			c++
			find(arr, []int{i, j + 1})
		}
	}
	if i > 0 && j > 0 { // check left up
		val := arr[i-1][j-1]
		if val != 0 && (val < 18 || val > 50) {

			c++
			find(arr, []int{i - 1, j - 1})
		}
	}
	if i > 0 && j <= 2 { // check right up
		val := arr[i-1][j+1]
		if val != 0 && (val < 18 || val > 50) {

			c++
			find(arr, []int{i - 1, j + 1})
		}
	}
	if i <= 3 && j <= 2 { // check right down
		val := arr[i+1][j+1]
		if val != 0 && (val < 18 || val > 50) {

			c++
			find(arr, []int{i + 1, j + 1})
		}
	}
	if i <= 3 && j > 0 { //check left down
		val := arr[i+1][j-1]
		if val != 0 && (val < 18 || val > 50) {

			c++
			find(arr, []int{i + 1, j - 1})
		}
	}
	return

}
