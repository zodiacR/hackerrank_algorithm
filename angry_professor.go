/*
Problem Statement
The professor is conducting a course on Discrete Mathematics to a class of N students. He is angry at the lack of their discipline, and he decides to cancel the class if there are less than K students present after the class starts.

Given the arrival time of each student, your task is to find out if the class gets cancelled or not.

Input Format
The first line of the input contains T, the number of test cases. Each test case contains two lines. The first line of each test case contains two space-separated integers, N and K.
The next line contains N space-separated integers indicating the arrival time of each student.

If the arrival time of a given student is a non-positive integer (N≤0), then the student enters before the class starts. If the arrival time of a given student is a positive integer (N>0), the student enters after the class has started. If a student enters the class exactly when it starts (N=0), the student is considered to have entered before the class has started.

Constraints
1<=T<=10
1<=N<=1000
1<=K<=N
The absolute value of arrival time does not exceed 100.

Output Format

For each testcase, print YES if the class gets cancelled and NO otherwise.

Sample Input

2
4 3
-1 -3 4 2
4 2
0 -1 2 1
Sample Output

YES
NO
*/
package main

import "fmt"

func main() {
	var cases, n, k int
	fmt.Scan(&cases)

	for i := 0; i < cases; i++ {
		fmt.Scan(&n)
		fmt.Scan(&k)

		count := 0
		arr := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scan(&arr[i])
			if arr[i] <= 0 {
				count++
			}
		}

		if count >= k {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
