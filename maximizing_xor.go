/*
Problem Statement
Given two integers, L and R, find the maximal values of A xor B, where A and B satisfy the following condition:
L<=A<=B<=R

Input Format
The input contains two lines; L is present in the first line and R in the second line.

Constraints
1<=L<=R<=103

Output Format
The maximal value as mentioned in the problem statement.

Sample Input
10
15

Sample Output
7
*/
package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	max := 0
	for i := a; i <= b; i++ {
		for j := a; j <= b; j++ {
			if m := i ^ j; m > max {
				max = m
			}
		}
	}

	fmt.Println(max)
}
