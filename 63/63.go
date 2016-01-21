/* The statement property can be formalized as follows:

We are looking for the 'i' for which we can find an 'n' and a 'k'  such that

	10^k == i^n
and
	n<=k<n+1

Since 10^2 is a 3-digit number, we have the condition i<10.

We can early out when i^n < 10^{n-1}.

Considering how small the input is, this can be solved by hand.

For 'i' in [1-9], the number of powers is

	n = floor(log(10)/(log(10)-log(i)))

Looping over powers requires handling big numbers. It is quite fast still.
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(^uint(0))
	result := 0
	for i := uint(1); i < 9; i++ {
		fmt.Println("::i=", i)
		power := i
		power10 := uint(1)

		for power10 <= power {
			fmt.Println(power10, "<=", power)
			result++
			power *= i
			power10 *= 10
		}
	}

	// Big numbers required for '9'.
	fmt.Println("::i=", 9)
	big9 := big.NewInt(9)
	big10 := big.NewInt(10)
	power := big.NewInt(9)
	power10 := big.NewInt(1)
	for power10.Cmp(power) <= 0 {
		fmt.Println(power, "<", power10)
		result++
		power.Mul(power, big9)
		power10.Mul(power10, big10)
	}

	fmt.Println(result)
}
