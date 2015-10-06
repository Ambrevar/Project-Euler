// Version with restricted input generation. Numbers are generated from the set
// {1,3,7,9}.

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

// Make sure that all rotations of 'tab' form prime numbers. Values of 'tab' are
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
		if sieve[n] {
			return false
		}
		n = (n%order)*10 + digits[tab[i]]
		if n == n_init {
			break
		}
	}

	return true
}

// Note that we are processing every rotation several times.
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
		(*array)[i] = 0
	}

	return true
}

var sieve []bool
var digits []uint

func main() {
	// Since each digit has to go in last position at some point, we can restrict
	// the input set.
	digits = []uint{1, 3, 7, 9}

	sieve = make_sieve(1000) // Limit is set to 1000000.

	array := make([]int8, 1, 6)
	array[0] = 3
	result := uint(0)
	for next_digitset(&array) {
		if primerot(array) {
			result++
		}
	}

	// Result is avove result + the four 1-digit prime numbers.
	// We process 1-digit primes separately since 2 and 5 are primes.
	fmt.Println(result + 4)
}
