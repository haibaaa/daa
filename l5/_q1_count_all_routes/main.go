/*
	memo dp
*/

package main

import (
	"fmt"
)

func countRoutes(locations []int, start, finish, fuel int) int {
	n := len(locations)

	// dp[city][remainingFuel] = number of ways to reach finish from city with remainingFuel

	// init
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, fuel+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(city, remainingFuel int) int
	dfs = func(city, remainingFuel int) int {
		if remainingFuel < 0 {
			return 0
		}

		if dp[city][remainingFuel] != -1 {
			return dp[city][remainingFuel]
		}

		// If we're at the finish city, count this as 1 route
		count := 0
		if city == finish {
			count = 1
		}

		// Try moving to all other cities
		for nextCity := 0; nextCity < n; nextCity++ {
			if nextCity == city {
				continue
			}

			fuelCost := abs(locations[city] - locations[nextCity])
			if remainingFuel >= fuelCost {
				count = (count + dfs(nextCity, remainingFuel-fuelCost))
			}
		}

		dp[city][remainingFuel] = count
		return count
	}

	return dfs(start, fuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, start, finish, fuel int
	fmt.Scan(&n, &start, &finish, &fuel)

	locations := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&locations[i])
	}

	result := countRoutes(locations, start, finish, fuel)
	fmt.Println(result)
}
