/*
Problem Statement
The Utopian Tree goes through 2 cycles of growth every year. The first growth cycle occurs during the spring, when it doubles in height. The second growth cycle occurs during the summer, when its height increases by 1 meter.
Now, a new Utopian Tree sapling is planted at the onset of spring. Its height is 1 meter. Can you find the height of the tree after N growth cycles?

Input Format
The first line contains an integer, T, the number of test cases.
T lines follow; each line contains an integer, N, that denotes the number of cycles for that test case.

Constraints
1<=T<=10
0<=N<=60

Output Format
For each test case, print the height of the Utopian Tree after N cycles. Each line thus has to contain a single integer, only.

Sample Input
3
0
1
4

Sample Output
1
2
7
*/

package main

import "fmt"

func main() {
	var n, yr int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&yr)
		h := 1

		for j := 0; j < yr; j++ {
			switch {
			case j%2 == 0:
				h *= 2
			case j%2 == 1:
				h++
			}
		}
		fmt.Println(h)
	}
}
