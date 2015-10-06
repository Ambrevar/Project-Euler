/* Brute force using home-made big integer multiplication. */

package main

import (
	"fmt"
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

func main() {
	result := 0
	array := make([]int, 1, 256)

	for a := 2; a < 100; a++ {
		array = array[:1]
		array[0] = 1

		for b := 2; b < 100; b++ {
			array = bigmul(array, a)
			sum := digitsum(array)
			if sum > result {
				result = sum
			}
		}
	}
	fmt.Println(result)
}
