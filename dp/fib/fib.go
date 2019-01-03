package main

import (
	"fmt"
)

/*
https://www.geeksforgeeks.org/?p=12635/   DP1
Dynamic Programming is an algorithmic paradigm that solves a given complex problem
by breaking it into subproblems and stores the results of subproblems to avoid
computing the same results again. Following are the two main properties of a problem
that suggests that the given problem can be solved using Dynamic Programming.
1. Overlapping Subproblems: when solutions of same subproblems are needed again and again.
In DP, computed solutions to subproblems are stored in a table so that these don't have to
be recomputed.
2. Optimal Substructure: A given problem has Optimal Substructure Property if
optimal solution of the given problem can be obtained by using optimal solutions
of its subproblems.
*/

func main() {
	fmt.Println(fib(10))
	fmt.Println(fibMemo(100))
	fmt.Println(fibTabulation(100))
	fmt.Println(fibSpaceOpt(100))
}

// Time complexity: O(2^N)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/*
The fib() method takes O(2^N) time complexity. It is not a practical approach.
To improve the algorithm, the results of subproblems can be stored for later use.
There are two ways to store the values:
1. Memorization(Top Down)
2. Tabulation(Bottom Up)
*/

// Top Down. Time: O(N), Space: O(N)
func fibMemo(n int) int {
	// Allocate memory to store the results of subproblems
	// Elements in Array are initialized to 0
	lookup := make([]int, n+1, n+1)
	return fibMemoHelper(n, lookup)
	//	return lookup[n]
}

func fibMemoHelper(n int, lookup []int) int {
	if n <= 1 {
		return n
	}

	if lookup[n] == 0 {
		lookup[n] = fibMemoHelper(n-1, lookup) + fibMemoHelper(n-2, lookup)
	}

	return lookup[n]
}

// Bottom Up. Time O(N), Space O(N)
func fibTabulation(n int) int {
	if n <= 1 {
		return n
	}

	lookup := make([]int, n+1, n+1)
	// By default, lookup[0] = 0
	lookup[1] = 1

	for i := 2; i <= n; i++ {
		lookup[i] = lookup[i-1] + lookup[i-2]
	}

	return lookup[n]
}

// For the bottom up approach, only the last two results need to be memorized.
// So, the space complexity can be improved to O(1)
func fibSpaceOpt(n int) int {
	if n <= 1 {
		return n
	}

	fib1, fib2 := 1, 0
	for i := 2; i <= n; i++ {
		//f := fib1 + fib2
		fib1, fib2 = fib1+fib2, fib1
	}

	return fib1
}
