/* Brute force: we loop over all possible primes.
This takes forever if we don't perform a few optimization:

- The `valid` function reduces the computation time tremendously by storing the result of
`isprime(concat(n1, n2)) && isprime(concat(n2, n1))` in an array.

- We break the loop when the current minimal sum is higher than the minimum sum.
*/

package main

import (
	"fmt"
	"math"
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

func concat(n1, n2 uint) uint64 {
	n2_bak := n2
	for n2 != 0 {
		n1 *= 10
		n2 = n2 / 10
	}
	return uint64(n1 + n2_bak)
}

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

func valid(n1, n2 uint, hash map[Pair]bool) bool {
	p := Pair{n1, n2}
	elem, ok := hash[p]
	if ok {
		return elem
	} else {
		elem = isprime(concat(n1, n2)) && isprime(concat(n2, n1))
		hash[p] = elem
		return elem
	}
}

type Pair struct {
	x, y uint
}

func main() {
	limit := uint(100)
	_, primes := make_sieve(limit)
	prime_count := len(primes) - 1
	hash := make(map[Pair]bool)

	// We use these constants to ease writing the stop condition for loops.
	// For instance, for some `i`, the current sum will be higher than 5*i + 2 + 4 + 6 + 8.
	const (
		size_i_f = 5
		size_i_o = 20
		size_j_f = 4
		size_j_o = 12
		size_k_f = 3
		size_k_o = 6
		size_l_f = 2
		size_l_o = 2
	)

	sum := uint(size_i_f*primes[prime_count] + size_i_o)
	for i := 1; i < prime_count && primes[i]*size_i_f+size_i_o < sum; i++ {
		limit_j := primes[i] + size_j_o
		for j := i + 1; j < prime_count && primes[j]*size_j_f+limit_j < sum; j++ {
			if valid(primes[i], primes[j], hash) {
				limit_k := primes[i] + primes[j] + size_k_o
				for k := j + 1; k < prime_count && primes[k]*size_k_f+limit_k < sum; k++ {
					if valid(primes[i], primes[k], hash) && valid(primes[j], primes[k], hash) {

						limit_l := primes[i] + primes[j] + primes[k] + size_l_o
						for l := k + 1; l < prime_count && primes[l]*size_l_f+limit_l < sum; l++ {

							if valid(primes[i], primes[l], hash) && valid(primes[j], primes[l], hash) && valid(primes[k], primes[l], hash) {

								limit_m := primes[i] + primes[j] + primes[k] + primes[l]
								for m := l + 1; m < prime_count && limit_m+primes[m] < sum; m++ {
									if valid(primes[i], primes[m], hash) && valid(primes[j], primes[m], hash) && valid(primes[k], primes[m], hash) && valid(primes[l], primes[m], hash) {
										new_sum := uint(primes[i] + primes[j] + primes[k] + primes[l] + primes[m])
										if new_sum < sum {
											// fmt.Println(primes[i], primes[j], primes[k], primes[l], primes[m], "Sum:", new_sum)
											sum = new_sum
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)
}
