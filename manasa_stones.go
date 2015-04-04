/*
Problem Statement
Manasa is out on a hike with friends. She finds a trail of stones with numbers on them. She starts following the trail and notices that two consecutive stones have a difference of either a or b. Legend has it that there is a treasure trove at the end of the trail and if Manasa can guess the value of the last stone, the treasure would be hers. Given that the number on the first stone was 0, find all the possible values for the number on the last stone.

Note: The numbers on the stones are in increasing order.

Input Format
The first line contains an integer T, i.e. the number of test cases. T test cases follow; each has 3 lines. The first line contains n (the number of stones). The second line contains a, and the third line contains b.

Output Format
Space-separated list of numbers which are the possible values of the last stone in increasing order.

Constraints
1<=T<=10
1<=n,a,b<=103

Sample Input
2
3
1
2
4
10
100

Sample Output
2 3 4
30 120 210 300
*/
package main

import "fmt"

func main() {
	var cases int
	fmt.Scan(&cases)

	for i := 0; i < cases; i++ {
		var stones, a, b int
		fmt.Scan(&stones, &a, &b)

		if stones == 0 {
			fmt.Println(0)
		} else if a == b {
			fmt.Printf("%d\n", a*(stones-1))
		} else {

			if a > b {
				a, b = b, a
			}

			for j := 0; j < stones; j++ {
				score := b*j + a*(stones-1-j)
				fmt.Printf("%d ", score)
			}
			fmt.Println()
		}
	}
}
