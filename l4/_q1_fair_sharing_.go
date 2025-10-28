package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// checks if coins can be divided fairly
func canDivideFairly(n int, coins []int) bool {
	// total sum of coins
	sum := 0
	// count of 1-rupee coins
	ones := 0
	// count of 2-rupee coins
	twos := 0
	for _, v := range coins {
		sum += v
		if v == 1 {
			// increment ones if coin is 1
			ones++
		} else {
			// increment twos otherwise (so it's 2)
			twos++
		}
	}
	// can't divide if sum is odd
	if sum%2 != 0 {
		return false
	}
	// amount each should get
	target := sum / 2
	// use as many 2-rupee coins as possible
	for twos > 0 && target >= 2 {
		target -= 2
		twos--
	}
	// check if rest can be made up with 1-rupee coins
	return target <= ones
}

func main() {
	// create scanner to read input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// read number of testcases
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		// read number of coins
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		// read line of coins as string
		line := scanner.Text()
		strCoins := strings.Split(line, " ")
		coins := make([]int, n)
		for j := 0; j < n; j++ {
			// parse each coin to int
			coins[j], _ = strconv.Atoi(strCoins[j])
		}
		// print yes if division is possible, else no
		if canDivideFairly(n, coins) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
