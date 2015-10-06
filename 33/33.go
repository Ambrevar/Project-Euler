/* Notation: "ab"/"bc" == (a*10+b)/(10*b+c)

When we have (a*10+b)/(10*b+c) == a/c, it means that

  a*(10*b-9*c) == c*b

It is easier and faster to check equality on integers than on floats.
*/

package main

import "fmt"

func simplify(a, b uint) (uint, uint) {
	if a > b {
		a, b = b, a
	}
	for i := uint(2); i*i <= a; i++ {
		for a%i == 0 && b%i == 0 {
			a /= i
			b /= i
		}
	}
	if a != 1 {
		if b%a == 0 {
			b /= a
			a = 1
		}
	}
	return a, b
}

func main() {
	prodnum := uint(1)
	prodden := uint(1)

	for a := uint(1); a < 10; a++ {
		// 'b' cannot be less than 'a' since the fraction is <1.
		for b := a; b < 10; b++ {
			c_start := uint(0)
			if a == b {
				if b == 9 {
					continue
				}
				// Make sure denominator > numerator.
				c_start = b + 1
			}
			for c := c_start; c < 10; c++ {
				if a*(10*b-9*c) == c*b {
					// fmt.Println(num, den)
					prodnum *= a
					prodden *= c
				}
			}
		}
	}

	_, y := simplify(prodnum, prodden)
	fmt.Println(y)
}
