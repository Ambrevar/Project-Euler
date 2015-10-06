/* When we find a non-Lychrel sequence, all numbers except the last are
non-Lychrel. However, it is not true for Lychrel sequences since the length of
the sequence varies depending on the starting number.

Considering that we can store all non-Lychrel numbers in an array to skip
numbers in the outer loop.

This is a bit faster.
*/

package main

import (
	"fmt"
	"math/big"
)

const NON_LYCHREL = 1

func reverse(n *big.Int) *big.Int {
	N := big.NewInt(0)
	mod := big.NewInt(0)
	result := big.NewInt(0)
	ten := big.NewInt(10)

	N.Set(n)
	for N.Cmp(big.NewInt(0)) > 0 {
		N.DivMod(N, ten, mod)
		result.Mul(result, ten)
		result.Add(result, mod)
	}
	return result
}

func palindromic(n *big.Int) bool {
	if n.Cmp(reverse(n)) == 0 {
		return true
	}
	return false
}

func main() {
	limit := 10000

	hit := [10001]int8{}

	bigLimit := big.NewInt(int64(limit))
	one := big.NewInt(1)
	sum := big.NewInt(0)

	result := 0

	for n := big.NewInt(1); n.Cmp(bigLimit) <= 0; n.Add(n, one) {
		if hit[n.Int64()] == NON_LYCHREL {
			result++
			continue
		}

		sum.Set(n)
		sequence := make([]*big.Int, 0, 100)
		for i := 1; i <= 50; i++ {
			r := reverse(sum)
			sequence = append(sequence, r)
			// We append sum to the sequence only here. Indeed, if it is palindromic
			// it does not mean it is not a Lychrel number.
			// Go warning: we cannot append 'sum' here since it is a pointer to a
			// value that we are going to modify right after.
			var s big.Int
			s.Set(sum)
			sequence = append(sequence, &s)

			sum.Add(sum, r)
			if palindromic(sum) {
				for _, v := range sequence {
					if v.Cmp(bigLimit) <= 0 {
						hit[v.Int64()] = NON_LYCHREL
					}
				}
				result++
				break
			}
		}
	}

	fmt.Println(10000 - result)
}
