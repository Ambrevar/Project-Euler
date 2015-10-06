package main

import "fmt"

// Check if 'n' can be part of a pandigital identity, i.e. if no digit is
// repeated. We return the table and the count to stack checks.
func pandigital_check(n uint, check *[10]bool, count *uint) (status bool) {
	status = true

	for ; n != 0; n /= 10 {
		*count++
		digit := n % 10
		if digit == 0 {
			status = false
			return
		}
		if !check[digit] {
			check[digit] = true
		} else {
			status = false
			return
		}
	}

	return
}

func main() {
	sum := uint(0)
	var unique [9876]bool

	// 'b' is the bigger of (a,b). The bounds below are because b would not be
	// pandigital otherwise.
	for b := uint(9876); b > 122; b-- {
		var check [10]bool
		count := uint(0)
		var status bool
		if status = pandigital_check(b, &check, &count); !status {
			continue
		}

		// 'a' is bound between values that make sure (a,b,a*b) is pandigital.
		var amin uint
		var amax uint
		if b > 999 {
			amin = 2
		} else {
			amin = 11
		}
		amax = 9876 / b

		for a := amin; a < amax; a++ {
			check_a := check
			count_a := count
			if status = pandigital_check(a, &check_a, &count_a); status {
				n := a * b
				if status = pandigital_check(n, &check_a, &count_a); status && count_a == 9 {
					// fmt.Println(b, a, n)
					if !unique[n] {
						sum += n
						unique[n] = true
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
