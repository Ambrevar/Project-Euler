// Count the number of identical cubes up to a permutation by using the sorted
// digits as a key in a map.
//
// Catch: when creating a key of digits in ascending order, zeros are not stored
// if the key is a number. Workarounds: either sort in descending order, or use
// strings, or prepend a digit in [1-9] to the key.
package main

import (
	"fmt"
	"sort"
)

var occur = map[uint]int{}
var roots = map[uint]uint{}
var results = []uint{}

func main() {
	limit := 5
	digits := 0

	for n := uint(1); ; n++ {
		n3 := n * n * n

		ref := []int{}
		for n3 != 0 {
			ref = append(ref, int(n3%10))
			n3 /= 10
		}

		if len(ref) > digits {
			// Next cubes have more digits and thus belong to a different class. Check
			// if any cube among the 'results' have the desired number of occurences,
			// and return the minimum.
			if len(results) > 1 {
				min := results[0]
				for _, key := range results {
					if occur[key] == limit && roots[key] < roots[min] {
						min = key
					}
				}
				if occur[min] == limit {
					fmt.Println(roots[min] * roots[min] * roots[min])
					return
				}
			}

			// Previous class did not yield any result. Free memory for the new class
			// of results.
			digits = len(ref)
			occur = map[uint]int{}
			roots = map[uint]uint{}
			results = []uint{}
		}

		sort.Ints(ref)
		// Prepend key with '1' to save zeros.
		key := uint(1)
		for _, v := range ref {
			key = 10*key + uint(v)
		}
		occur[key]++
		// Remember first root that yields the key.
		if _, ok := roots[key]; !ok {
			roots[key] = n
		}

		if occur[key] == limit {
			results = append(results, key)
		}

	}
}
