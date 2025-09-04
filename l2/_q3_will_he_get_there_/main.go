package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s1, s2 string
	if _, err := fmt.Fscan(in, &s1); err != nil {
		return
	}
	if _, err := fmt.Fscan(in, &s2); err != nil {
		return
	}

	// find destination
	target := 0
	for _, c := range s1 {
		if c == '+' {
			target++
		} else {
			target--
		}
	}

	// count '?'
	current, q := 0, 0
	for _, c := range s2 {
		switch c {
		case '+':
			current++
		case '-':
			current--
		case '?':
			q++
		}
	}

	delta := target - current
	var ways float64

	if q == 0 {
		ways = 1
	} else {
		// Check if a valid number of '+' assignments exists
		if (q+delta)%2 == 0 {
			r := (q + delta) / 2
			if r >= 0 && r <= q {
				ways = nCrFloat(q, r)
			}
		}
	}

	// Total possibilities = 2^q
	total := math.Pow(2, float64(q))
	prob := ways / total

	fmt.Printf("%.12f\n", prob)
}

func nCrFloat(n, r int) float64 {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}
	num := 1.0
	den := 1.0
	for i := 1; i <= r; i++ {
		num *= float64(n - r + i)
		den *= float64(i)
	}
	return num / den
}
