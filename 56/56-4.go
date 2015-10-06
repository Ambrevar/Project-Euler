/* Use bound of 56-3 on the home-made big multiplication from 56-1. */

package main

import (
	"fmt"
	"math"
)

func bigmul(array []int, n int) []int {
	carry := 0

	for k, v := range array {
		carry += v * n
		array[k] = carry % 10
		carry /= 10
	}
	for carry > 0 {
		array = append(array, carry%10)
		carry /= 10
	}
	return array
}

func digitsum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}

func max_digitsum(a, b int) int {
	a_float := float64(a)
	b_float := float64(b)
	return int(9 * (1 + math.Floor(b_float*math.Log10(a_float))))
}

func min_power(a, max int) int {
	a_float := float64(a)
	max_float := float64(max)

	return int((max_float/9 - 2) / math.Log10(a_float))
}

func main() {
	limit := 100
	result := 0
	array := make([]int, 1, 256)

	for a := limit - 1; a >= 0; a-- {
		if result > max_digitsum(a, limit-1) {
			fmt.Println(result)
			return
		}

		array = array[:1]
		array[0] = 1
		min := min_power(a, result)

		for b := 1; b < min; b++ {
			array = bigmul(array, a)
		}

		for b := min; b < limit; b++ {
			array = bigmul(array, a)
			sum := digitsum(array)
			if sum > result {
				result = sum
			}
		}
	}

}
