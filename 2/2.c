/* A sum of different numbers below N will not exceed N*(N+1)/2. We store the
values in 64 bits numbers to avoid overflow.

The explicit formula is slower since it involves a square root, multiplications
and divisions.

Optimization trick: every 3 term is even, so we sum thrice per loop. At the end,
we need to subtract each of the last 2 numbers that were above the limit, if
any.
*/

#include <stdio.h>
#include <stdint.h>

int main() {
	uint64_t limit = 4000000;

	uint64_t result = 2;
	uint64_t fibo1 = 1;
	uint64_t fibo2 = 2;
	uint64_t buffer;

	while (fibo1 + fibo2 <= limit) {
		buffer = fibo2;
		fibo2 = fibo1 + fibo2;
		fibo1 = buffer;

		buffer = fibo2;
		fibo2 = fibo1 + fibo2;
		fibo1 = buffer;

		buffer = fibo2;
		fibo2 = fibo1 + fibo2;
		fibo1 = buffer;

		result += fibo2;
	}

	if (fibo2 > limit) {
		result -= fibo2;
	}

	if (fibo1 > limit) {
		result -= fibo1;
	}

	printf("%ld\n", result);
	return 0;
}
