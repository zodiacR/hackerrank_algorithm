/*
Problem Statement
Little Bob loves chocolate, and he goes to a store with $N in his pocket. The price of each chocolate is $C. The store offers a discount: for every M wrappers he gives to the store, he gets one chocolate for free. How many chocolates does Bob get to eat?

Input Format:
The first line contains the number of test cases, T.
T lines follow, each of which contains three integers, N, C, and M.

Output Format:
Print the total number of chocolates Bob eats.

Constraints:
1<=T<=1000 2<=N<=105
1<=C<=N
2<=M<=N

Sample input
3
10 2 5
12 4 4
6 2 2

Sample Output
6
3
5
*/
package main

import "fmt"

func main() {
	var t, n, c, m int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n, &c, &m)

		amount := n / c
		wrapper := amount

		for {
			if wrapper >= m {
				cho := wrapper / m
				amount += cho
				wrapper = cho + wrapper - cho*m
			} else {
				fmt.Println(amount)
				break
			}
		}
	}
}
