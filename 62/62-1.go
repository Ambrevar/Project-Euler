// Start from N, check if N+1...P are permutations.
// P is the last number for which P^3 has the same number of digit than N^3.
// Memoize roots for a minor speed-up.
//
// This is slow.
//
// Besides, the problem statement asks for exactly 5 occurences, while here we
// include occurences >= 5. See the alternative verison for a fix.
package main

import (
	"fmt"
)

func isPermutation(ref []int, n uint) bool {
	for n != 0 {
		ref[n%10]--
		n /= 10
	}
	for _, v := range ref {
		if v != 0 {
			return false
		}
	}
	return true
}

func int2array(n uint) []int {
	ref := make([]int, 10)
	for n != 0 {
		ref[n%10]++
		n /= 10
	}

	return ref
}

var memoize = map[uint]int{}

func main() {
	limit := 5

	for n := uint(1); ; n++ {
		if memoize[n] {
			continue
		}
		count := 1
		n3 := n * n * n
		n3_10 := 10 * n3
		for p := n + 1; ; p++ {
			ref := int2array(n3)
			p3 := p * p * p
			if p3 > n3_10 {
				break
			}
			if isPermutation(ref, p3) {
				memoize[p] = true
				count++
				if count == limit {
					fmt.Println(n3)
					return
				}
			}
		}
	}
}
