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
	b := []int{10, 22, 9, 90, 33, 21, 50, 60, 2, 80, 81, 82, 85}
	fmt.Println(lisOSIntuitive(a), lisOSIntuitive(b))
	fmt.Println(lisOSTopDown(a), lisOSTopDown(b))
	fmt.Println(lisBottomUp(a), lisBottomUp(b))
	fmt.Println(lis(a), lis(b))
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
	_, len := lisOSIntuitiveHelper(a)
	return len
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
// Returns first value is the max LIS of current element
//         second value is the max LIS of all elements in [0...n-1]
func lisOSIntuitiveHelper(a []int) (int, int) {
	n := len(a)
	if n == 1 {
		// The first element
		return 1, 1
	}

	maxLen := 1
	maxLenSofar := 1
	for i := 1; i < n; i++ {
		len, lenSofar := lisOSIntuitiveHelper(a[:i])
		maxLenSofar = max(maxLenSofar, lenSofar)
		// The slicing doesn't include i, so the last element is at i-1
		if a[n-1] > a[i-1] && maxLen < len+1 {
			maxLen = len + 1
		}
	}

	return maxLen, max(maxLenSofar, maxLen)
}

// Memorization. Top Down
// Time: O(N^2), space: O(N)
func lisOSTopDown(a []int) int {
	n := len(a)
	lis := make([]int, n, n)
	lisOSTopDownHelper(a, lis)

	//fmt.Println(lis)
	return maxInArray(lis)
}

// A little difference from the intuitive method.
// There is no maxLenSofar returned because it has difficulty to get it.
func lisOSTopDownHelper(a []int, lis []int) int {
	n := len(a)
	if lis[n-1] > 0 {
		return lis[n-1]
	}

	maxLen := 1
	for i := 1; i < n; i++ {
		len := lisOSTopDownHelper(a[:i], lis)
		if a[n-1] > a[i-1] && maxLen < len+1 {
			maxLen = len + 1
		}
	}

	// Store result of the subproblem
	lis[n-1] = maxLen

	return maxLen
}

// Tabulation approach. Bottom Up
// Time: O(N^2), space: O(N)
func lisBottomUp(a []int) int {
	n := len(a)
	lis := make([]int, n, n)
	lis[0] = 1

	for i := 1; i < n; i++ {
		lis[i] = 1 // Init value
		for j := 0; j < i; j++ {
			if a[j] < a[i] && lis[i] < lis[j]+1 {
				// Store solution of subproblem
				lis[i] = lis[j] + 1
			}
		}
	}

	//fmt.Println(lis)
	return maxInArray(lis)
}

// Find maximum value in an array
func maxInArray(a []int) int {
	maxLen := a[0]
	for i := 1; i < len(a); i++ {
		maxLen = max(maxLen, a[i])
	}

	return maxLen
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
