// We compute powers over the last 10 digits since the other digits do not
// reflect on the result.
package main

import "fmt"

func power(n uint64) uint64 {
	p := n
	for i := uint64(1); i < n; i++ {
		p = p * n % 10000000000
	}
	return p
}

func main() {
	limit := uint64(1000)
	result := uint64(0)
	for i := uint64(1); i <= limit; i++ {
		result += power(i)
	}

	fmt.Println(result % 10000000000)
}
