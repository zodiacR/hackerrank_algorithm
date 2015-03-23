/*
Problem Statement
You are given a square map of size n×n. Each cell of the map has a value denoting its depth. We will call a cell of the map a cavity if and only if this cell is not on the border of the map and each cell adjacent to it has strictly smaller depth. Two cells are adjacent if they have a common side.

You need to find all the cavities on the map and depict them with the uppercase character X.

Input Format
The first line contains an integer, n, denoting the size of the map. Each of the following n lines contains n positive digits without spaces. Each digit (1-9) denotes the depth of the appropriate area.

Constraints
1≤n≤100

Output Format
Output n lines, denoting the resulting map. Each cavity should be replaced with character X.

Sample Input
4
1112
1912
1892
1234

Sample Output
1112
1X12
18X2
1234
*/
package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	Map := make([][]rune, n)
	MapCopy := make([][]rune, n)

	var row string

	for i := 0; i < n; i++ {
		Map[i] = make([]rune, n)
		MapCopy[i] = make([]rune, n)
		fmt.Scan(&row)

		for j := 0; j < n; j++ {
			Map[i][j] = rune(row[j])
			MapCopy[i][j] = rune(row[j])
		}
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			gap1 := Map[i][j] - Map[i][j-1]
			gap2 := Map[i][j] - Map[i-1][j]
			gap3 := Map[i][j] - Map[i+1][j]
			gap4 := Map[i][j] - Map[i][j+1]

			if gap1 > 0 && gap2 > 0 && gap3 > 0 && gap4 > 0 {
				MapCopy[i][j] = rune('X')
			}

		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%c", MapCopy[i][j])
		}
		fmt.Println()
	}
}
