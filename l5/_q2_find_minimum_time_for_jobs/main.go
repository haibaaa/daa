/*
 * bitmask dp 2^n states
 */
package main

import (
	"fmt"
	"math"
)

func minimizeMaxWorkTime(jobs []int, k int) int {
	n := len(jobs)

	// dp[mask] = minimum max working time when jobs in mask are assigned
	// We'll use a different approach: try all assignments with backtracking
	workers := make([]int, k)
	result := math.MaxInt32

	// Sort jobs in descending order for better pruning
	sortDesc(jobs)

	var backtrack func(jobIdx int)
	backtrack = func(jobIdx int) {
		if jobIdx == n {
			// All jobs assigned, find max working time
			maxTime := 0
			for _, time := range workers {
				if time > maxTime {
					maxTime = time
				}
			}
			if maxTime < result {
				result = maxTime
			}
			return
		}

		// Pruning: if current max already >= result, no point continuing
		currentMax := 0
		for _, time := range workers {
			if time > currentMax {
				currentMax = time
			}
		}
		if currentMax >= result {
			return
		}

		// Try assigning current job to each worker
		seen := make(map[int]bool)
		for i := 0; i < k; i++ {
			// Pruning: skip if we've already tried this workload
			if seen[workers[i]] {
				continue
			}
			seen[workers[i]] = true

			workers[i] += jobs[jobIdx]
			backtrack(jobIdx + 1)
			workers[i] -= jobs[jobIdx]
		}
	}

	backtrack(0)
	return result
}

func sortDesc(arr []int) {
	// Simple bubble sort in descending order
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	jobs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&jobs[i])
	}

	result := minimizeMaxWorkTime(jobs, k)
	fmt.Println(result)
}
