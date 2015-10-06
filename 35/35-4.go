// Version with restricted input generation. Numbers are generated from the set
// {1,3,7,9}. We check with a bool slice if number has not already been
// processed from a rotation.

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

func make_number(tab []int8) (n uint, length uint, order uint) {
	n = 0
	length = 0
	order = 1
	for _, v := range tab {
		n = n*10 + digits[v]
		order *= 10
		length++
	}
	order /= 10
	return
}

// Make sure that all rotations of 'tab' form prime numbers. Values of 'tab' are
// indices of 'digits'. We do not count duplicates, so we stop as soon as a
// rotated number is equal to the first generated number.
func primerot(n uint, length uint, order uint) (bool, uint) {
	n_init := n
	count := uint(0)
	for i := uint(0); i < length; i++ {
		check[n] = true
		count++
		if sieve[n] {
			return false, 0
		}
		n = (n%order)*10 + n/order
		if n == n_init {
			break
		}
	}

	return true, count
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
var check []bool

func main() {
	// Since each digit has to go in last position at some point, we can only
	// use these.
	digits = []uint{1, 3, 7, 9}
	check = make([]bool, 1000000)

	sieve = make_sieve(1000) // Limit is set to 1000000.

	array := make([]int8, 1, 6)
	array[0] = 3
	result := uint(0)
	for next_digitset(&array) {
		n, length, order := make_number(array)
		if !check[n] {
			if status, res := primerot(n, length, order); status {
				result += res
			}
		}
	}

	// Result is avove result + the four 1-digit prime numbers.
	// We process 1-digit primes separately since 2 and 5 are primes.
	fmt.Println(result + 4)
}
