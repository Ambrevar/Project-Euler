// Faster version. TODO: Explain this code.
// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
// The list of values for the fraction are given by `(root + addend) / denominator`.
// The first value is `root` and the last value is `root + addend`
package main

import (
	"fmt"
)

func main() {
	root := 1
	oddCount := 0
	for n := 2; n < 10000; n++ {
		if (root+1)*(root+1) == n {
			root++
		} else {
			denominator := 1
			addend := 0
			period := 0

			for {
				addend = -(addend - ((root+addend)/denominator)*denominator)
				denominator = (n - addend*addend) / denominator
				period++
				if denominator == 1 {
					break
				}
			}

			if period%2 == 1 {
				oddCount++
			}
		}
	}

	fmt.Println(oddCount)
}
