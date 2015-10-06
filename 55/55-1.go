/* Brute-force version. */

package main

import (
	"fmt"
	"math/big"
)

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

	bigLimit := big.NewInt(int64(limit))
	one := big.NewInt(1)
	sum := big.NewInt(0)

	result := 0

	for n := big.NewInt(1); n.Cmp(bigLimit) <= 0; n.Add(n, one) {
		sum.Set(n)
		for i := 1; i <= 50; i++ {
			sum.Add(sum, reverse(sum))
			if palindromic(sum) {
				result++
				break
			}
		}
	}

	fmt.Println(10000 - result)
}
