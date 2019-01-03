package main

import (
	"fmt"
)

/*
https://www.geeksforgeeks.org/longest-increasing-subsequence-dp-3/
Longest Increasing Subsequence DP3
The LIS problem is to find the length of the longest subsequence of a given sequence
such that all elments of the subsequence are sorted in increasing order. For example,
the length of LIS for {10, 22, 9, 33, 21, 50, 41, 60, 80} is 6 and LIS is {10, 22, 33, 50, 60, 80}
*/

func main() {
	a := []int{10, 22, 9, 33, 21, 50, 41, 60, 80, 41, 60, 9}
	fmt.Println(lisOSIntuitive(a))
	fmt.Println(lisOSTopDown(a))
	fmt.Println(lisBottomUp(a))
	fmt.Println(lis(a))
}

/*
Optimal Substructure:
Let a[0...n-1] be the input array and L(i) be the length of the LIS ending at index i
such that a[i] is the last element of the LIS.
Then, L(i) can be recursively written as:
L(i) = 1 + max(L(j)) where 0 < j < i and a[j] < a[i]; or
L(i) = 1, if no such j exists.
To find the LIS for a given array, we need to return max(L(i)) where 0 < i < n.
Thus, we see the LIS problem satisfies the optimal substructure property as the main
problem can be solved using solutions to subproblems
*/
func lisOSIntuitive(a []int) int {
	return lisOSIntuitiveHelper(a)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Time complexity: exponential
// For lis(n) -> lis(n-1), lis(n-2), ..., lis(1)
//     lis(n-1) -> lis(n-2), ..., lis(1)
func lisOSIntuitiveHelper(a []int) int {
	length := len(a)
	if length == 1 {
		// The first element
		return 1
	}

	maxLen := 1
	maxValue := a[0]
	for i := 1; i < length; i++ {
		len := lisOSIntuitiveHelper(a[:i])
		// The slicing doesn't include i, so the last element is at i-1
		if a[length-1] > maxValue {
			len++
		}
		maxValue = max(maxValue, a[i])
		maxLen = max(maxLen, len)
	}

	return maxLen
}

// Memorization. Top Down
// Time: O(N^2), space: O(N)
func lisOSTopDown(a []int) int {
	lookup := make([]int, len(a), len(a))
	maxLen := lisOSTopDownHelper(a, lookup)
	//fmt.Println(lookup)
	return maxLen
}

func lisOSTopDownHelper(a []int, lookup []int) int {
	length := len(a)
	if lookup[length-1] > 0 {
		return lookup[length-1]
	}

	maxLen := 1
	maxValue := a[0]
	for i := 1; i < length; i++ {
		len := lisOSTopDownHelper(a[:i], lookup)
		if a[length-1] > maxValue {
			len++
		}
		maxLen = max(maxLen, len)
		maxValue = max(maxValue, a[i])
	}

	// Store result of the subproblem
	lookup[length-1] = maxLen

	return maxLen
}

// Tabulation approach. Bottom Up
// Time: O(N^2), space: O(N)
func lisBottomUp(a []int) int {
	len := len(a)
	lis := make([]int, len, len)
	lis[0] = 1

	for i := 1; i < len; i++ {
		// By default, set current length of LIS to the value of previous element
		lis[i] = lis[i-1]
		maxValue := a[0]
		for j := 0; j < i; j++ {
			maxValue = max(maxValue, a[j])
			if maxValue < a[i] && lis[i] < lis[j]+1 {
				// Store solution of subproblem
				lis[i] = lis[j] + 1
			}
		}
	}

	//fmt.Println(lis)
	return lis[len-1]
}

/*
O(NlogN) approach
https://www.geeksforgeeks.org/longest-monotonically-increasing-subsequence-size-n-log-n/
Create a new array to store the increasing sequence.
Take array {10, 22, 9, 33, 21, 50, 41, 60, 80} as an example.
First, add 10 to the result array, [10].
Then next element 22 is greater than 10, append to the list. [10, 22]
Next element 9 is less than the last one, replace 10. [9, 22]
Next 33 > 22, append to the list. [9, 22, 33]
Next 21, replace 22. [9, 21, 33]
Next 50, append to the list. [9, 21, 33, 50]
and so on...
Finally return the length of the result array
Since the result array is sorted, the ceiling value can be found by binary search,
The time complexity is O(NlogN)
*/
func lis(a []int) int {
	res := make([]int, 0, len(a))
	res = append(res, a[0])

	for i := 1; i < len(a); i++ {
		if a[i] > res[len(res)-1] {
			// Append to the end of result array
			res = append(res, a[i])
		} else {
			// Replace
			ceiling := findCeiling(res, a[i])
			res[ceiling] = a[i]
		}
		//fmt.Println(res)
	}
	return len(res)
}

// Binary search to find the ceiling of the give key
func findCeiling(a []int, key int) int {
	low, high := 0, len(a)-1

	for low < high {
		mid := low + (high-low)/2
		if key > a[mid] {
			low = mid + 1 // + 1 is must, other with it results in endless loop
		} else {
			high = mid
		}
	}
	return low
}
