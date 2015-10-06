/* Similar to 60-1.
We use Miller-Rabin as primality test.

Store the order of the the primes to save some cycles every time we check if a
pair is concat-prime.

Let's we call 'i' the outer loop variable, the inner loops go over 2..i instead of i..limit.
Going this way has two advantages:

- We do the concat-prime tests only for 'i', as the others have been done during
  the former iterations of i.
- The first result is the right result.

Despite some useless looping, this is the fastest method.
*/

package main

import (
	"fmt"
)

// See problem 3.
func make_sieve(limit uint64) (sieve []bool, primes []uint64) {
	var bound uint64
	if limit > 1 {
		bound = limit * limit
	} else if limit == 1 {
		bound = 2
	} else {
		bound = 1
	}

	sieve = make([]bool, bound)
	primes = make([]uint64, 0, bound)
	sieve[0] = true
	if limit == 0 {
		return
	}
	sieve[1] = true
	primes = primes[:len(primes)+1]
	primes[0] = 1
	for i := uint64(2); i <= limit; i++ {
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
	const limit = 100
	const prime_limit = 2000
	_, primes := make_sieve(limit)
	prime_count := len(primes) - 1

	// Entry hash[x][y] must be so that x>y.
	// A static array takes some space but saves a lot of time over a map.
	hash := [prime_limit][prime_limit]bool{}

	i_order := uint64(10)
	// We start looping from the seond prime since the first prime '2' cannot be
	// part of a family.
	for i := 2; i < prime_count; i++ {
		if primes[i] > i_order {
			i_order *= 10
		}

		// Fill in the hash.
		j_order := uint64(10)
		for j := 2; j < i; j++ {
			if primes[j] > j_order {
				j_order *= 10
			}
			if miller_rabin(primes[i]*j_order+primes[j]) && miller_rabin(primes[j]*i_order+primes[i]) {
				hash[i][j] = true
			}
		}

		// Try out the families.
		for j := 2; j < i; j++ {
			if !hash[i][j] {
				continue
			}

			for k := 2; k < j; k++ {
				if !hash[i][k] || !hash[j][k] {
					continue
				}

				for l := 2; l < k; l++ {
					if !hash[i][l] || !hash[j][l] || !hash[k][l] {
						continue
					}

					for m := 2; m < l; m++ {

						if !hash[i][m] || !hash[j][m] || !hash[k][m] || !hash[l][m] {
							continue
						}

						fmt.Println(primes[i] + primes[j] + primes[k] + primes[l] + primes[m])
						return
					}
				}
			}
		}
	}
}
