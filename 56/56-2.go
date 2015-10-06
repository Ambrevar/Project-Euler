/* Using standard library instead of home made big integer support.
This is slower because of the required conversion to decimal in 'digitsum'.
*/

package main

import (
	"fmt"
	"math/big"
)

// This is too slow.
func digitsum(n *big.Int) int {
	sum := 0
	text, _ := n.MarshalText()
	for _, v := range text {
		sum += int(v - '0')
	}
	return sum
}

func main() {
	result := 0

	for a := int64(2); a < 100; a++ {
		v := big.NewInt(a)
		mul := big.NewInt(a)

		for b := 2; b < 100; b++ {
			mul.Mul(mul, v)
			sum := digitsum(mul)

			if sum > result {
				result = sum
			}
		}
	}

	fmt.Println(result)
}
