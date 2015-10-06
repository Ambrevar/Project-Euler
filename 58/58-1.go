// Because of the dimension of the problem, a sieve is much to slow and eats to
// much memory. We need trial division or Miller-Rabbin.
// See problem 28.

package main

import (
	"fmt"
	"math"
)

func isprime(n uint64) bool {
	if n == 2 || n == 3 || n == 5 {
		return true
	}
	if n <= 1 || n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return false
	}

	root := uint64(math.Sqrt(float64(n)))
	step := uint64(2)
	for i := uint64(7); i < root; i += step {
		if n%i == 0 {
			return false
		}
		step = 6 - step
	}
	return true
}

func main() {
	diag_count := 5.0
	diag_primes := 3.0
	n := uint64(9)
	step := uint64(4)

	for diag_primes/diag_count >= 0.1 {
		for i := 1; i <= 3; i++ {
			n += step
			if isprime(n) {
				diag_primes++
			}
		}
		n += step
		step += 2
		diag_count += 4
	}

	fmt.Println(step - 1)
}
