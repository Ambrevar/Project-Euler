/* In 60-3 we go through useless loop cycles.

Let's consider a concat-prime family 'j, k, l' with no m < l such that 'j, k, l,
m' is a concat-prime family. For every i > j, we will go through the same dead
end with this 'j, k, l' tuple.

Top optimize this we build the concat-prime families recursively.

Despite the loop saving, the bookkeeping is too heavy and makes the agorithm
slower overall.
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

/* Return all the concat-prime families of cardinal 'cardinal' and made of 'n'
and numbers < n.

'n' is not stored in the result.

We save the results in the 'families' structure.

If cardinal is 0, return an empty family.
If cardinal is 1, return singleton families.
*/
func derive(n int, order uint64, cardinal int) [][]Pair {
	result := make([][]Pair, 0)

	if cardinal <= 0 {
		return result

	} else if cardinal == 1 {
		// No need to browse subfamilies, they are empty. We create families with
		// all integers < n that are concat-primes with n.

		j_order := uint64(10)
		for j := 2; j < n; j++ {
			if primes[j] > j_order {
				j_order *= 10
			}

			if !miller_rabin(primes[n]*j_order+primes[j]) || !miller_rabin(primes[j]*order+primes[n]) {
				continue
			}

			result = append(result, make([]Pair, 1))
			result[len(result)-1][0] = Pair{j, j_order}
		}

		families[cardinal][n] = result
		return result
	}

	// Cardinal >= 2.
	j_order := uint64(10)
	for j := 2; j < n; j++ {
		if primes[j] > j_order {
			j_order *= 10
		}

		if !miller_rabin(primes[n]*j_order+primes[j]) || !miller_rabin(primes[j]*order+primes[n]) {
			continue
		}

		if families[cardinal-1][j] == nil {
			families[cardinal-1][j] = derive(j, j_order, cardinal-1)
		}

		for _, member := range families[cardinal-1][j] {
			valid := true

			for _, v := range member {
				if !miller_rabin(uint64(primes[n]*v.order+primes[v.n])) || !miller_rabin(uint64(primes[v.n]*order+primes[n])) {
					valid = false
					break
				}
			}

			if valid {
				subarray := make([]Pair, len(member)+1)
				copy(subarray, member)
				subarray[len(member)] = Pair{j, j_order}
				result = append(result, subarray)
			}
		}
	}

	families[cardinal][n] = result
	return result
}

// We use the pair structure to save the number.
type Pair struct {
	n     int
	order uint64
}

const limit = 100
const prime_limit = 2000
const cardinal_limit = 5

/* families[cardinal][n] stores the list of concat-prime families of cardinal
`cardinal` for n. It means that every member of the families is less than n.

n is the n-th prime number.

For instance, 7 has the following families of cardinal 0 and 1:

  families[0][5] = []    // Undefined, no families.
  families[1][5] = [[2]] // 37 and 73 are primes.

From the problem statement, we have:

  families[3][122] = [[2, 4, 29]]
*/
var families [cardinal_limit][prime_limit][][]Pair
var primes []uint64

func main() {
	_, primes = make_sieve(limit)
	prime_count := len(primes) - 1

	var result [][]Pair

	// We start looping from 2 since the first prime '2' cannot be part of a family.
	i_order := uint64(10)
	for i := 2; i < prime_count; i++ {
		if primes[i] > i_order {
			i_order *= 10
		}
		result = derive(i, i_order, cardinal_limit-1)
		if len(result) != 0 {
			for _, j := range result {
				sum := uint64(0)
				for _, k := range j {
					sum += primes[k.n]
				}
				fmt.Println(sum + primes[i])
			}
			return
		}
	}
}
