package main

import (
	"fmt"
)

func main() {
	a1 := []int{3, 4, 9, 1}
	a2 := []int{5, 3, 8, 9, 10, 4, 2, 1}
	fmt.Println(lcisRecursive(a1, a2))

	len, seq := lcisRecursiveWithPath(a1, a2)
	fmt.Println("length:", len, ", LCIS:", seq)

	fmt.Println(lcis(a1, a2))
	fmt.Println(lcisPath(a1, a2))
}

/* https://www.geeksforgeeks.org/longest-common-increasing-subsequence-lcs-lis/
Given two arrays, find length of the longest common increasing subsequence (LCIS)
and print one of such sequence(multiple sequences may exist)
Suppose we have two arrays:
a1[] = {3, 4, 9, 1}
a2[] = {5, 3, 8, 9, 10, 2, 1}
Our answer would be {3, 9} and the LCIS is 2
*/

// Exponential time complexity?
func lcisRecursive(a1 []int, a2 []int) int {
	if len(a1) == 0 || len(a2) == 0 {
		return 0
	}

	maxLcis := 0
	for i := 0; i < len(a1); i++ {
		lcisSofar := 0
		for j := 0; j < len(a2); j++ {
			if a1[i] == a2[j] {
				// Recursively do it from new start
				lcisSofar = 1 + lcisRecursiveHelper(a1[i+1:], a2[j+1:], a1[i])
			}
		}
		maxLcis = max(maxLcis, lcisSofar)
	}

	return maxLcis
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcisRecursiveHelper(a1, a2 []int, x int) int {
	if len(a1) == 0 || len(a2) == 0 {
		return 0
	}

	for i := 0; i < len(a1); i++ {
		if x < a1[i] {
			for j := 0; j < len(a2); j++ {
				if a1[i] == a2[j] {
					return 1 + lcisRecursiveHelper(a1[i+1:], a2[j+1:], a1[i])
				}
			}
		}
	}

	return 0
}

// First return value is the length of LCIS
// Second return value is the subarray of LCIS
// Only return one subsequence
func lcisRecursiveWithPath(a1, a2 []int) (int, []int) {
	if len(a1) == 0 || len(a2) == 0 {
		return 0, nil
	}

	maxLen := 0
	var seq []int
	for i := 0; i < len(a1); i++ {
		lenCurrent := 0
		var seqCurrent []int
		for j := 0; j < len(a2); j++ {
			if a1[i] == a2[j] {
				seqCurrent = append(seqCurrent, a1[i])
				length, newSeq := lcisRecursiveWithPathHelper(a1[i+1:], a2[j+1:], a1[i])
				lenCurrent = length + 1
				seqCurrent = append(seqCurrent, newSeq...)
			}
		}

		if lenCurrent > maxLen {
			maxLen = lenCurrent
			seq = seqCurrent
		}
	}

	return maxLen, seq
}

func lcisRecursiveWithPathHelper(a1, a2 []int, x int) (int, []int) {
	if len(a1) == 0 || len(a2) == 0 {
		return 0, nil
	}

	var seq []int
	for i := 0; i < len(a1); i++ {
		if x < a1[i] {
			for j := 0; j < len(a2); j++ {
				if a1[i] == a2[j] {
					seq = append(seq, a1[i])
					length, newSeq := lcisRecursiveWithPathHelper(a1[i+1:], a2[j+1:], a1[i])
					return length + 1, append(seq, newSeq...)
				}
			}
		}
	}

	return 0, nil
}

// Time: O(M*N), space: O(N)
func lcis(a1, a2 []int) int {
	if len(a1) == 0 || len(a2) == 0 {
		return 0
	}

	// Allocate an array to store the LCIS in a2
	table := make([]int, len(a2))

	for i := 0; i < len(a1); i++ {
		current := 0
		for j := 0; j < len(a2); j++ {
			// If both arrays has the same element
			if a1[i] == a2[j] && current+1 > table[j] {
				table[j] = current + 1
			}

			// Now seek for previous smaller common element for current element in a1
			if a1[i] > a2[j] && table[j] > current {
				current = table[j]
			}
		}
		//fmt.Println(table)
	}

	maxLen := table[0]
	for i := 1; i < len(table); i++ {
		maxLen = max(maxLen, table[i])
	}

	//fmt.Println(table)

	return maxLen
}

func lcisPath(a1, a2 []int) int {
	if len(a1) == 0 || len(a2) == 0 {
		return 0
	}

	// Allocate an array to store the LCIS in a2
	table := make([]int, len(a2))
	parent := make([]int, len(a2))

	for i := 0; i < len(a1); i++ {
		current := 0
		// This index means the start of LCIS
		last := -1
		for j := 0; j < len(a2); j++ {
			// If both arrays has the same element
			if a1[i] == a2[j] && current+1 > table[j] {
				table[j] = current + 1
				parent[j] = last
				last = j
			}

			// Now seek for previous smaller common element for current element in a1
			if a1[i] > a2[j] && table[j] > current {
				current = table[j]
				last = j
			}
		}
		//fmt.Println("In loop LCIS", table)
		//fmt.Println("In loop parent", parent)
	}

	// Find the max length
	var sq []int
	sq = append(sq, 0)
	maxLen := 0
	for i := 1; i < len(table); i++ {
		if maxLen < table[i] {
			maxLen = table[i]
			sq = append(sq[:0], i)
		} else if maxLen == table[i] {
			sq = append(sq, i)
		}
	}

	//fmt.Println(table)
	//fmt.Println(parent)

	// Print LCIS
	for i := 0; i < len(sq); i++ {
		path := make([]int, maxLen)
		for index, j := sq[i], maxLen-1; index != -1; index = parent[index] {
			path[j] = a2[index]
			j--
		}
		fmt.Println(path)
	}

	return maxLen
}
