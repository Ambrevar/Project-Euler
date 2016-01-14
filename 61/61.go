/* Store all 4-digit polygonal numbers into tables, except octagonal numbers.
Index them by their first 2 digits so to speed up cycle creation.

To find cycles, we loop over octagonal numbers since they have the
widest gap between two consecutive numbers.

Find cycle over the all the permutations of {heptagonals, ..., triangles}.

Early out on first complete cycle since we know there is only one.

The problem statement mentions that the polygonal types should be represented by
different numbers. It is uncleear if it means 'a different element of the set',
or 'all elements must have a different value'.

*/
package main

import (
	"fmt"
)

// Return 0 when no cycle is found.
func sumCycle(set [][100][]int, n int, remainder int, order int) int {
	polygons := set[n]
	if n == len(set)-1 {
		for _, p := range polygons[remainder] {
			if p%100 == order {
				return p
			}
		}
		return 0
	}

	for _, p := range polygons[remainder] {
		if next := sumCycle(set, n+1, p%100, order); next != 0 {
			return next + p
		}
	}

	return 0
}

// Call 'sumCycle' over all permutations of 'set'.
func permute(set [][100][]int, n int, remainder int, order int) int {
	if n == len(set)-1 {
		if res := sumCycle(set, 0, remainder, order); res != 0 {
			return res
		}
		return 0
	}

	for i := n; i < len(set); i++ {
		set[n], set[i] = set[i], set[n]
		if res := permute(set, n+1, remainder, order); res != 0 {
			return res
		}
		set[n], set[i] = set[i], set[n]
	}
	return 0
}

func main() {
	triangle := [100][]int{}
	for n := 1; ; n++ {
		p := n * (n + 1) / 2
		if p < 1000 {
			continue
		}
		if p > 9999 {
			break
		}
		index := p / 100
		triangle[index] = append(triangle[index], p)
	}

	square := [100][]int{}
	for n := 1; ; n++ {
		p := n * n
		index := p / 100
		if p < 1000 || index < 10 {
			continue
		}
		if p > 9999 {
			break
		}
		square[index] = append(square[index], p)
	}

	pentagonal := [100][]int{}
	for n := 1; ; n++ {
		p := n * (3*n - 1) / 2
		index := p / 100
		if p < 1000 || index < 10 {
			continue
		}
		if p > 9999 {
			break
		}
		pentagonal[index] = append(pentagonal[index], p)
	}

	hexagonal := [100][]int{}
	for n := 1; ; n++ {
		p := n * (2*n - 1)
		index := p / 100
		if p < 1000 || index < 10 {
			continue
		}
		if p > 9999 {
			break
		}
		hexagonal[index] = append(hexagonal[index], p)
	}

	heptagonal := [100][]int{}
	for n := 1; ; n++ {
		p := n * (5*n - 3) / 2
		index := p / 100
		if p < 1000 || index < 10 {
			continue
		}
		if p > 9999 {
			break
		}
		heptagonal[index] = append(heptagonal[index], p)
	}

	set := [][100][]int{heptagonal, hexagonal, pentagonal, square, triangle}

	// Find cycles starting from octagonal numbers.
	for n := 1; ; n++ {
		p := n * (3*n - 2)
		if p < 1000 {
			continue
		}
		if p > 9999 {
			break
		}

		result := permute(set, 0, p%100, p/100)
		if result != 0 {
			fmt.Println(result + p)
			return
		}
	}
}
