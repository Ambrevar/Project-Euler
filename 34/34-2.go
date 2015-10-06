// Use the same technique as problem 30.
package main

import "fmt"

// Precompute the factorials.
var factorial = [10]uint{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func next_digitset(array *[]int8) bool {
	i := 0
	// Go to first digit that is not a 9.
	for i < len(*array) && (*array)[i] == 9 {
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

/* WARNING: `array` integrity is not guaranteed. You should copy the array
before calling this function. */
func curious(array []int8) (uint, bool) {
	fsum := uint(0)
	for _, i := range array {
		fsum += factorial[i]
	}

	result := fsum

	for fsum > 0 {
		n := int8(fsum % 10)
		old_fsum := fsum
		for k, v := range array {
			if n == v {
				array[k] = array[len(array)-1]
				array = array[:len(array)-1]
				fsum /= 10
				break
			}
		}
		if fsum == old_fsum {
			return 0, false
		}
	}
	if len(array) > 0 {
		return 0, false
	}
	return result, true
}

func main() {
	sum := uint(0)
	// 7*9! == 2540160 < 9999999 (7 digits), so number higher than 2540160 can
	// never be curious.

	// We traverse all numbers with a different digit set from 2 digits to 7 digits.
	array := make([]int8, 2, 7)
	// Start from 10.
	array[0] = 9

	// Use a unique buffer to minimize memory allocation.
	buffer := make([]int8, 2, 7)

	for next_digitset(&array) {
		buffer = buffer[:len(array)]
		copy(buffer, array)
		if s, status := curious(buffer); status {
			sum += s
		}
	}
	fmt.Println(sum)
}
