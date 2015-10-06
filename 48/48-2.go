/* We divide the computation into 3 chunks: 2 chunks of 4 digits and the last
chunk of 2. The size of the last chunk does not matter since result will be
truncated past the 10th digit. The chunk must be able to contain the last power
(here 1000).

We use the descending fast exponentiation algorithm. We loop from the most
significant bit of 'power' to the last. On each bit we square the value. Each
time we hit a 1, we multiply the value by 'power'.

The reason we do not use the ascending algorithm is that we would need another
big number to store the power.
*/

package main

import (
	"fmt"
)

func main() {
	chunksize := 10000

	var power, k int
	var (
		result0, result1, result2 int
		mul0, mul1, mul2          int
		temp0, temp1, temp2       int
	)

	for power = 1; power <= 1000; power++ {
		// This sets mul = 1.
		mul0 = 1
		mul1 = 0
		mul2 = 0
		for k = 1; k <= power; k <<= 1 {
			// Move 'k' to power's last bit.
		}
		k >>= 1
		for ; k > 0; k >>= 1 {
			// This sets mul *= mul.
			// It is like pen & paper multiplication.
			temp0 = mul0 * mul0
			temp1 = 2*mul0*mul1 + temp0/chunksize
			temp2 = 2*mul0*mul2 + mul1*mul1 + temp1/chunksize
			mul0 = temp0 % chunksize
			mul1 = temp1 % chunksize
			mul2 = temp2 % chunksize

			if power&k != 0 {
				// This sets mul *= power.
				temp0 = mul0 * power
				temp1 = mul1*power + temp0/chunksize
				temp2 = mul2*power + temp1/chunksize
				mul0 = temp0 % chunksize
				mul1 = temp1 % chunksize
				mul2 = temp2 % chunksize
			}
		}

		// This sets result += mul.
		temp0 = mul0 + result0
		temp1 = mul1 + result1 + temp0/chunksize
		temp2 = mul2 + result2 + temp1/chunksize
		result0 = temp0 % chunksize
		result1 = temp1 % chunksize
		result2 = temp2 % chunksize
	}

	fmt.Printf("%02v%04v%04v\n", result2%100, result1, result0)
}
