// Using home-made big-digit operations. Much faster thanks to instant length
// computation.
package main

import (
	"fmt"
)

// Write a+b to a.
func bigadd(a, b []byte) []byte {
	carry := byte(0)

	// Extend a if it cannot hold b.
	if len(a) < len(b) {
		buf := make([]byte, len(b))
		copy(buf, a)
		a = buf
	}

	for i := range a {
		sum := a[i] + b[i] + carry
		carry = sum / 10
		a[i] = sum - carry*10
	}
	if carry != 0 {
		a = append(a, carry)
	}
	return a
}

func main() {
	limit := 1000
	result := 0

	d := make([]byte, 1)
	n := make([]byte, 1)
	d[0] = 1
	n[0] = 1
	for i := 1; i <= limit; i++ {
		buf := make([]byte, len(d))
		copy(buf, d)
		d = bigadd(d, n)
		n = bigadd(buf, d)

		if len(n)-len(d) > 0 {
			result++
		}
	}

	fmt.Println(result)
	return
}
