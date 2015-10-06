/* Smarter version, more complex, but quite faster.

Starting from prime 3, we check next prime P that concat with it.
We intersect the list of concat primes.
Repeat with P until we have a family of 4 or limit is reached.
If limit is reached, backtrack and check next prime.

The solution family is:
  {13, 5197, 5701, 6733, 8389}
Sum: 26033
*/

package main

import (
	"fmt"
	"math"
	"os"
)

// At limit = 136, first solution is not the right one. limit = 100 gives the
// right result faster.
const limit = 100

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

type Family struct {
	index        uint
	intersection map[uint]bool
	pos          uint
	sum          uint
}

func next(family *Family, table [][]uint) bool {
	for i := family.pos + 1; i < uint(len(table[family.index])); i++ {
		v := table[family.index][i]
		if family.intersection[v] {
			family.pos = i
			return true
		}
	}
	return false
}

/* `index` is only used if `family` is empty. */
func add(family []Family, table [][]uint, primes []uint, index uint) []Family {
	last := len(family)
	family = family[:last+1]

	if last == 0 {
		family[last].index = uint(index)
	} else {
		family[last].index = table[family[last-1].index][family[last-1].pos]
	}

	// Init intersection here, because we return in table creation if it is empty,
	// and we need a valid intersection in the caller.
	family[last].intersection = make(map[uint]bool)

	// Create intersection and table if required.
	for table[family[last].index] == nil {
		for j := uint(family[last].index + 1); j < uint(len(primes)); j++ {
			if isprime(concat(primes[family[last].index], primes[j])) && isprime(concat(primes[j], primes[family[last].index])) {
				table[family[last].index] = append(table[family[last].index], j)
			}
		}

		// In case we found a prime that has no concat primes, we retry with a
		// higher index. This is only useful for family[0].
		if table[family[last].index] == nil {
			if last == 0 {
				family[last].index++
				if family[last].index >= uint(len(table)) {
					return family
				}
			} else {
				return family
			}
		}
	}

	family[last].sum = primes[family[last].index]
	if last != 0 {
		family[last].sum += family[last-1].sum
	}

	// Note: could be faster if we move the comparison 'last==0' outside. This
	// induce code duplcation however, and in practice there is almost no
	// difference.
	for _, v := range table[family[last].index] {
		if last == 0 || family[last-1].intersection[v] {
			family[last].intersection[v] = true
		}
	}

	// Position on first table element that is in intersection.
	family[last].pos = 0
	if !family[last].intersection[table[family[last].index][0]] {
		next(&family[last], table)
	}

	return family
}

/* table[p] = <list of concat prime>

The p-th element is the p-th prime number. The `table` holds all prime number
indices within range that have the concatenation property.

We keep track of the current family through the
`family` table.

	family[i] = i-th member of the family, primes[i] is the real value.

The intersection of the concat prime lists from family member 0 to i is kept in
the hash family[i].intersection.
*/
func main() {
	_, primes := make_sieve(limit)
	table := make([][]uint, len(primes))

	// For a set of N concat primes, we need size = N-1.
	const size = 4
	family := make([]Family, 0, size)
	// Init with 2nd prime `3` (primes[2] = 3): we ignore `2`, since it is never a concat prime.
	family = add(family, table, primes, uint(2))

	sum := uint(0)
	for i := 0; i <= size; i++ {
		sum += primes[len(primes)-1-i]
	}

	for {
		if len(family) > 0 {
			if len(family[len(family)-1].intersection) != 0 && family[len(family)-1].sum < sum {
				if len(family) == size {
					pos := table[family[len(family)-1].index][family[len(family)-1].pos]
					new_sum := family[len(family)-1].sum + primes[pos]

					if new_sum < sum {
						sum = new_sum
					}

					// WARNING: Breaking here prints the first result only. This might not
					// be the right one depending on the value of `limit`. Remove the
					// break to compute the correct value independently of the limit.
					break

					// Force Backtrack
					family[len(family)-1].intersection = nil
				} else {
					family = add(family, table, primes, 0)
				}
			} else {
				// Backtrack.
				family = family[:len(family)-1]

				// Find next position in previous element. If not found, remove it again.
				for len(family) > 1 && !next(&family[len(family)-1], table) {
					family = family[:len(family)-1]
				}

				// Special case if first element of the family: we increment the index.
				if len(family) == 1 && !next(&family[len(family)-1], table) {
					if family[0].index >= uint(len(primes)-1) {
						fmt.Fprintln(os.Stderr, "Result not found within range", primes[family[0].index])
						break
					}
					new_index := family[0].index + 1
					family = family[:len(family)-1]
					family = add(family, table, primes, new_index)
				}
			}
		} else {
			break
		}
	}
	fmt.Println(sum)
}
