/* We need big numbers since the numerator and the denominator will grow very big.

This is a little slow because of the conversion to decimal.
*/

package main

import (
	"fmt"
	"math/big"
)

// Return digit_count(n)-digit_count(d).
func digit_count_cmp(n, d *big.Int) int {
	n_text, _ := n.MarshalText()
	d_text, _ := d.MarshalText()

	if len(n_text) > len(d_text) {
		return 1
	}
	return 0
}

var zero *big.Int
var ten *big.Int

func main() {
	limit := 1000
	result := 0

	zero = big.NewInt(0)
	ten = big.NewInt(10)
	n := big.NewInt(1)
	d := big.NewInt(1)
	buf := big.NewInt(0)

	for i := 1; i <= limit; i++ {
		buf.Set(d)
		d.Add(d, n)
		n.Add(d, buf)
		if digit_count_cmp(n, d) > 0 {
			result++
		}
	}

	fmt.Println(result)
	return
}
