/*
Problem Statement
There are N integers in an array A. All but one integer occur in pairs. Your task is to find the number that occurs only once.

Input Format
The first line of the input contains an integer N, indicating the number of integers. The next line contains N space-separated integers that form the array A.

Constraints
1<=N<100
N % 2=1 (N is an odd number)
0<=A[i]<=100,∀i∈[1,N]
Output Format

Output S, the number that occurs only once.

Sample Input:1
1
1

Sample Output:1
1

Sample Input:2
3
1 1 2

Sample Output:2
2

Sample Input:3
5
0 0 1 2 1

Sample Output:3
2
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)
	stdin.ReadLine()
	raw, _, _ := stdin.ReadLine()
	nums := strings.Split(string(raw[:len(raw)]), " ")

	numbers := make(map[string]int)

	for _, v := range nums {
		if _, ok := numbers[v]; ok {
			numbers[v]++
		} else {
			numbers[v] = 1
		}
	}

	fmt.Printf("%v\n", numbers)

	for k, value := range numbers {
		if value == 1 {
			fmt.Println(k)
			break
		}
	}
}
