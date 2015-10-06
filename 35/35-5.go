/* Version with restricted input generation. Numbers are generated from the set
 {1,3,7,9}.

We generate only ascending numbers, from which we derive permutations. Since we
do not rotations, we permutate all digits but one.

For instance, 971 rotates to 197 and 719. From 971 we derive 917, which rotates
to 179 and 791. We do not need to permutate 9.

The input set is thus slightly reduced, but the additional maintainance cost
counterbalances this gain.
*/

package main

import "fmt"

// See problem 3.
func make_sieve(limit uint) (sieve []bool) {
	var bound uint
	if limit > 1 {
		bound = limit * limit
	} else if limit == 1 {
		bound = 2
	} else {
		bound = 1
	}

	sieve = make([]bool, bound)
	sieve[0] = true
	if limit == 0 {
		return
	}
	sieve[1] = true
	for i := uint(2); i <= limit; i++ {
		if !sieve[i] {
			for j := i * i; j < bound; j += i {
				sieve[j] = true
			}
		}
	}
	return
}

// Make sure that all rotations of 'tab' give prime numbers. Values of 'tab' are
// indices of 'digits'. We do not count duplicates, so we stop as soon as a
// rotated number is equal to the first generated number.
func primerot(tab []int8) bool {
	n := uint(0)

	order := uint(1)
	for _, v := range tab {
		n = n*10 + digits[v]
		order *= 10
	}
	order /= 10

	n_init := n

	limit := len(tab)
	for i := 0; i < limit; i++ {
		if sieve[n] == true || perm_hit[n] == true {
			return false
		}
		perm_hit[n] = true
		n = (n%order)*10 + digits[tab[i]]
		if n == n_init {
			break
		}
	}

	return true
}

// Perumatations might hit duplicates (e.g. 311 can permutate in 311). We check
// for hits before.
func perm_rec_sub(array []int8, sub []int8, count *uint) {
	if len(sub) == 1 {
		n := uint(0)
		for _, v := range array {
			n = 10*n + uint(v)
		}
		if primerot(array) == true {
			*count += uint(len(array))
		}
	} else {
		for k := range sub {
			sub[k], sub[0] = sub[0], sub[k]
			perm_rec_sub(array, sub[1:], count)
			sub[k], sub[0] = sub[0], sub[k]
		}
	}
}

// See problem 30.
func next_digitset(array *[]int8) bool {
	i := 0
	// Go to first digit that is not a 3.
	for i < len(*array) && (*array)[i] == 3 {
		i++
	}

	if i == cap(*array) {
		// No more permutation.
		return false
	}
	if i >= len(*array) {
		*array = (*array)[:len(*array)+1]
	} else {
		// We are still using the same number of digits. Set all previous digits to
		// the value of this one to avoid duplicate digit sets.
		(*array)[i]++
	}
	i--
	for ; i >= 0; i-- {
		(*array)[i] = (*array)[i+1]
	}

	return true
}

var sieve []bool
var digits []uint
var perm_hit []bool

func main() {
	// Since each digit has to go in last position at some point, we can restrict
	// the input set.
	digits = []uint{1, 3, 7, 9}

	sieve = make_sieve(1000) // Limit is set to 1000000.

	// This hit table avoids counting duplicates in permutations (i.e. 311). 11 is
	// still counted twice.
	perm_hit = make([]bool, 1000000)

	array := make([]int8, 1, 6)
	array[0] = 3
	result := uint(0)
	for next_digitset(&array) {
		perm_rec_sub(array, array[1:], &result)
	}

	// Result is avove result + the four 1-digit prime numbers.
	// We process 1-digit primes separately since 2 and 5 are primes.
	// We subtract 1 since 11 is counted twice because of permutations.
	fmt.Println(result + 4 - 1)
}
