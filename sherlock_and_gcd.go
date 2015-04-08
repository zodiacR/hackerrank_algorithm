/*
Problem Statement

Sherlock is stuck while solving a problem: Given an array A={a1,a2,⋯,aN}, he wants to know if there exists a subset B of this array which follows these statements:

B is a non-empty subset.
There exists no integer x(x>1) which divides all elements of B.
There are no elements of B which are equal to another.
Input Format

The first line of input contains an integer, T, representing the number of test cases. Then T test cases follow.
Each test case consists of two lines. The first line contains an integer, N, representing the size of array A. In the second line there are N space-separated integers, a1,a2,…,an, representing the elements of array A.

Constraints
1<=T<=10
1<=N<=100
1<=ai<=105 ∀1<=i<=N

Output Format
Print YES if such a subset exists; otherwise, print NO.

Sample Input
3
3
1 2 3
2
2 4
3
5 5 5

Sample Output
YES
NO
NO
*/
package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	return a
}

func main() {
	var t, n int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		var a, b int

		fmt.Scan(&a)
		for j := 1; j < n; j++ {
			fmt.Scan(&b)
			a = GCD(a, b)
		}
		if a == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
