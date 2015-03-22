/*
Problem Statement
You are given an integer N. Find the digits in this number that exactly divide N (division that leaves 0 as remainder) and display their count. For N=24, there are 2 digits (2 & 4). Both of these digits exactly divide 24. So our answer is 2.

Note
If the same number is repeated twice at different positions, it should be counted twice, e.g., For N=122, 2 divides 122 exactly and occurs at ones' and tens' position. So for this case, our answer is 3.
Division by 0 is undefined.

Input Format
The first line contains T (the number of test cases), followed by T lines (each containing an integer N).

Constraints
1<=T<=15
0<N<1010

Output Format
For each test case, display the count of digits in N that exactly divide N in a separate line.

Sample Input
2
12
1012

Sample Output
2
3
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)

	var num string

	for i := 0; i < n; i++ {
		fmt.Scan(&num)

		dec := strconv.Atoi(num)
		count := 0

		for _, v := range num {
			i := int(v - '0')
			if i != 0 {
				div := dec / i
				if r := i * div; r == dec {
					count++
				}
			}
		}

		fmt.Println(count)
	}
}
