/* Miller-Rabin primality test is deterministic under some conditions.
See http://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test.

It is much faster than trial division.
*/

package main

import (
	"fmt"
)

// Fast exponentiation x^y % mod.
func pow_mod(x, y, mod uint64) uint64 {
	power := x
	result := uint64(1)

	for y > 0 {
		if y%2 == 1 {
			result = (result * power) % mod
		}
		power = (power * power) % mod
		y /= 2
	}
	return result
}

// Return true if composite.
func witness(n, a uint64) bool {
	// Write n as d * 2^s, with d odd.
	d := n / 2
	s := 1
	for d%2 == 0 {
		d /= 2
		s++
	}

	power_a := pow_mod(a, d, n)
	if power_a == 1 {
		return false
	}
	for i := 1; i <= s; i++ {
		if power_a == n-1 {
			return false
		}
		power_a = (power_a * power_a) % n
	}

	return true
}

// Testing n < 4,759,123,141 is determistic with witnesses 2, 7 and 61.
func miller_rabin(n uint64) bool {
	if witness(n, 2) || witness(n, 7) || witness(n, 61) {
		return false
	}

	return true
}

func main() {
	diag_count := 5.0
	prime_count := 3.0
	n := uint64(9)
	step := uint64(4)

	for prime_count/diag_count >= 0.1 {
		for i := 1; i <= 3; i++ {
			n += step
			if miller_rabin(n) {
				prime_count++
			}
		}
		n += step
		step += 2
		diag_count += 4
	}

	fmt.Println(step - 1)
}
