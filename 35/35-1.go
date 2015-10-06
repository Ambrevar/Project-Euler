// Brute force version.

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

// Make sure that all rotations of 'n' form prime numbers. We do not count
// duplicates, so we stop as soon as a rotated number is equal to the first
// generated number.
func primerot(n uint) (status bool, count uint) {
	count = 0
	status = true

	order := uint(1)
	digits := 0
	for nb := n; nb != 0; nb /= 10 {
		order *= 10
		digits++
	}
	order /= 10

	n_init := n

	for i := 0; i < digits; i++ {
		check[n] = true
		if sieve[n] {
			status = false
			continue
		}
		count++
		n = (n%order)*10 + n/order
		if n == n_init {
			break
		}
	}

	return
}

var sieve []bool

// Check if number has already been processed.
var check []bool

func main() {
	sieve = make_sieve(1000) // Limit is set to 1000000.
	check = make([]bool, 1000000)

	result := uint(0)
	for i := uint(10); i < 1000000; i++ {
		if !check[i] {
			if status, count := primerot(i); status {
				result += count
			}
		}
	}

	// Result is avove result + the four 1-digit prime numbers.
	fmt.Println(result + 4)
}
