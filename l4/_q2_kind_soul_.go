package main

import (
	"fmt"
	"sort"
)

// generates all distinct partitions of n using numbers <= max
func distinctPartitions(n int, max int, curr []int, results *[][]int) {
	// if sum is zero, store current partition
	if n == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*results = append(*results, tmp)
		return
	}
	// iterate from min(max, n)
	for i := min(max, n); i >= 1; i-- {
		// ensure distinct numbers by decreasing order
		if len(curr) == 0 || i < curr[len(curr)-1] {
			distinctPartitions(n-i, i-1, append(curr, i), results)
		}
	}
}

// returns smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// read input number n
	var n int
	fmt.Scan(&n)

	// generate all unique partitions of n into distinct parts
	results := [][]int{}
	distinctPartitions(n, n, []int{}, &results)

	// print number of partitions found
	fmt.Println(len(results))

	// sort each partition in descending order and print
	for _, r := range results {
		sort.Slice(r, func(i, j int) bool { return r[i] > r[j] })
		for _, num := range r {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
