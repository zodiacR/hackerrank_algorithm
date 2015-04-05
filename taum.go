/*
Problem Statement
Taum is planning to celebrate the birthday of his friend, Diksha. There are two types of gifts that Diksha wants from Taum: one is black and the other is white. To make her happy, Taum has to buy B number of black gifts and W number of white gifts.

The cost of each black gift is X units.
The cost of every white gift is Y units.
The cost of converting each black gift into white gift or vice versa is Z units.
Help Taum by deducing the minimum amount he needs to spend on Diksha's gifts.

Input Format
The first line will contain an integer T which will be the number of test cases.
There will be T pairs of lines. The first line of each test case will contain the values of integers B and W. Another line of each test case will contain the values of integers X, Y, and Z.

Constraints
1<=T<=10
0<=X,Y,Z,B,W<=109

Output Format
T lines, each containing an integer: the minimum amount of units Taum needs to spend on gifts.

Sample Input
5
10 10
1 1 1
5 9
2 3 4
3 6
9 1 1
7 7
4 2 1
3 3
1 9 2

Sample Output
20
37
12
35
12
*/

package main

import "fmt"

func main() {
	var t, x, y, z, b, w int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&b, &w, &x, &y, &z)
		var result int

		if x > y {
			su := y + z
			if x > su {
				result = y*(b+w) + z*b
			} else {
				result = x*b + y*w
			}
		} else {
			su := x + z
			if y > su {
				result = x*(b+w) + z*w
			} else {
				result = x*b + y*w
			}
		}

		fmt.Println(result)
	}
}
