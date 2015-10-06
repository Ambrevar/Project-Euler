/* We look for the '1' digits found in the number but the last digit. They act
as a mask. Actually any digit but '0' would work since the first digit cannot be
a 0. Then we rotate the '1' with all possible digits and check for primality.
*/

package main

import (
	"fmt"
)

// See problem 3.
func make_sieve(limit uint) (sieve []bool, primes []uint) {
	var bound uint
	if limit > 1 {
		bound = limit * limit
	} else if limit == 1 {
		bound = 2
	} else {
		bound = 1
	}

	sieve = make([]bool, bound)
	primes = make([]uint, 0, bound)
	sieve[0] = true
	if limit == 0 {
		return
	}
	sieve[1] = true
	primes = primes[:len(primes)+1]
	primes[0] = 1
	for i := uint(2); i <= limit; i++ {
		if !sieve[i] {
			primes = primes[:len(primes)+1]
			primes[len(primes)-1] = i
			for j := i * i; j < bound; j += i {
				sieve[j] = true
			}
		}
	}

	// If we would only care about the first limit-th values, this loop would
	// not be necessary.

	for i := limit + 1; i < bound; i++ {
		if !sieve[i] {
			primes = primes[:len(primes)+1]
			primes[len(primes)-1] = i
		}
	}

	return
}

func main() {
	sieve, primes := make_sieve(1000)

	/* Possible optimization: with less than 6 digits, 2, 4 or 5 substitutions
	will generate at least 3 numbers divisible by 3, so the family cardinal cannot
	be more than 7. */
	for _, n := range primes {
		buf := n / 10
		offset := uint(0)
		order := uint(10)

		// We keep the most significant digit in 'buf' to see where to start the
		// trials, since the most significant digit cannot be 0.
		for ; buf > 9; buf /= 10 {
			if buf%10 == 1 {
				offset = offset + 1*order
			}
			order *= 10
		}

		var start uint
		if buf == 1 {
			start = n
			offset = offset + 1*order
		} else {
			start = n - offset
		}

		if offset == 0 {
			// No '1' was found.
			continue
		}

		// We add 8 times offset since we chose '1' for the mask, and the max value
		// is obviously '9'.
		limit := n + 8*offset
		cardinal := 0
		for i := start; i <= limit; i += offset {
			if !sieve[i] {
				cardinal++
			}
		}
		if cardinal >= 8 {
			fmt.Println(start)
			break
		}
	}
}
