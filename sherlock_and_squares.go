/*
Problem Statement
Watson gives two integers (A and B) to Sherlock and asks if he can count the number of square integers between A and B (both inclusive).

Note: A square integer is an integer which is the square of any integer. For example, 1, 4, 9, and 16 are some of the square integers as they are squares of 1, 2, 3, and 4, respectively.

Input Format
The first line contains T, the number of test cases. T test cases follow, each in a new line.
Each test case contains two space-separated integers denoting A and B.

Output Format
For each test case, print the required answer in a new line.

Constraints
1<=T<=100
1<=A<=B<=109

Sample Input
2
3 9
17 24

Sample output
2
0
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var t int

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b int64
		fmt.Scan(&a, &b)

		b_sqrt := int64(math.Sqrt(float64(b)))
		a_sqrt := int64(math.Sqrt(float64(a)))

		count := b_sqrt - a_sqrt
		if math.Sqrt(float64(a))-float64(a_sqrt) == 0 {
			count++
		}
		fmt.Println(count)
	}
}
