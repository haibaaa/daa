package main

import (
	"fmt"
)

/*create a subsetSum array*/
/*do bs over all possible answer*/
/*check if each possible answer is achieveable*/

var subsetSum []int
var memo map[int]bool

func canPart(mask int, k int, limit int) bool {
	if mask == 0 {
		return true
	}
	if k == 0 {
		return false
	}

	key := mask<<4 | k // because k<2^4
	if val, exists := memo[key]; exists {
		return val
	}

	for sub := mask; sub > 0; sub = (sub - 1) & mask {
		if subsetSum[sub] <= limit {
			if canPart(mask^sub, k-1, limit) {
				memo[key] = true
				return true
			}
		}
	}

	memo[key] = false
	return false
}

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	totalMasks := 1 << n
	subsetSum = make([]int, totalMasks)

	for mask := 0; mask < totalMasks; mask++ {
		sum := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += jobs[i]
			}
		}
		subsetSum[mask] = sum
	}

	/*binary search*/
	low := 0
	for i := 0; i < n; i++ {
		if jobs[i] > low {
			low = jobs[i]
		}
	}
	high := 0
	for i := 0; i < n; i++ {
		high += jobs[i]
	}

	for low < high {
		mid := (low + high) / 2
		memo = make(map[int]bool)

		if canPart(totalMasks-1, k, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	jobs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&jobs[i])
	}

	result := minimumTimeRequired(jobs, k)
	fmt.Println(result)
}
