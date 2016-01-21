package main

import (
	"fmt"
)

type expansion struct {
	m int
	d int
}

func continuedFracPeriod(s int, root int) int {
	// Period does not take the floored square root into account.
	period := -1

	periodic := map[expansion]bool{}
	m := 0
	d := 1
	a0 := root
	a := root
	// Perfect squares are a special case and have a period of 0.
	if a*a == s {
		return 0
	}

	for _, ok := periodic[expansion{m: m, d: d}]; !ok; _, ok = periodic[expansion{m: m, d: d}] {
		period++
		periodic[expansion{m: m, d: d}] = true
		m = d*a - m
		d = (s - m*m) / d
		a = (a0 + m) / d
	}
	fmt.Println("s:", s, "d:", d, "root:", a0, "a:", a)
	return period
}

func main() {
	limit := 10000
	result := 0
	root := 1
	for i := 2; i <= limit; i++ {
		if (root+1)*(root+1) < i {
			root++
		}
		if continuedFracPeriod(i, root)%2 == 1 {
			result++
		}
	}
	fmt.Println(result)
}
