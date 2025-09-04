package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func countPairs(arr []int, ceil int) int {
	i, j := 0, len(arr)-1
	count := 0
	for i < j {
		if (arr[i] + arr[j]) <= ceil {
			count += j - i
			i++
		} else {
			j--
		}
	}

	return count
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var arr []int
		var n, l, r int
		fmt.Fscan(in, &n, &l, &r)

		for ; n > 0; n-- {
			var x int
			fmt.Fscan(in, &x)
			arr = append(arr, x)
		}

		// sort to use two pointer
		sort.Ints(arr)

		ans := countPairs(arr, r) - countPairs(arr, l-1)
		fmt.Println(ans)
	}
}
