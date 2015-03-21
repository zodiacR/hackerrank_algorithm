package main

import "fmt"

func main() {
	var n, t int

	fmt.Scan(&n, &t)

	units := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&units[i])
	}

	var i, j int
	for k := 0; k < t; k++ {
		fmt.Scan(&i, &j)

		min := units[i]
		for s := i; s <= j; s++ {
			if min > units[s] {
				min = units[s]
			}
		}

		fmt.Println(min)
	}
}
