package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(arr []int) int {
	var sum int
	for _, elt := range arr {
		sum += elt
	}
	tg := sum / 2

	dp := make([]bool, tg+1)
	dp[0] = true
	for _, elt := range arr {
		for i := tg; i >= elt; i-- {
			dp[i] = dp[i] || dp[i-elt]
		}
	}

	for i := tg; i >= 0; i-- {
		if dp[i] {
			return (sum - 2*i) + 1
		}
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)
	var arr []int

	for ; t > 0; t-- {
		var x int
		fmt.Fscan(in, &x)
		arr = append(arr, x)
	}

	fmt.Println(solve(arr))
}
