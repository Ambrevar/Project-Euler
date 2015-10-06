// Use the same technique as problem 30.
package main

import "fmt"

// Precompute the factorials.
var factorial = [10]uint{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func curious(n uint) bool {
	n_old := n
	result := uint(0)
	for n != 0 {
		result += factorial[n%10]
		n /= 10
	}
	if n_old == result {
		return true
	}
	return false
}

func main() {
	sum := uint(0)
	// 7*9! == 2540160 < 9999999 (7 digits), so number higher than 2540160 can
	// never be curious.
	for i := uint(10); i < 2540160; i++ {
		if curious(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
