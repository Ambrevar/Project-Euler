/* Trial division is expensive for big numbers with a lot of prime divisors, but
in this special case it happens to be the fastest.

See the other implementations for something more generic.

Details: we divide the input by incrementing numbers starting from 2. This way
we only use prime numbers in the divisions.

Starting from 7, prime numbers are found within the steps 4, 2, 4, 2, ... We
could find some more complicated pattern to diminish the step frequency, but
then the step computation would take too much time. The '4-2' sequence is the
most efficient.

We can stop at the square root: if i*i>n, and n is not divisible by any prime
below i, then n is either 1 or prime.
*/

#include <stdio.h>
#include <stdint.h>

int main() {
	uint64_t limit = 600851475143;
	uint64_t i;

	/* Factors < 7. */
	unsigned int init[3] = {2, 3, 5};
	for (i = 0; i < 3; i++) {
		while (limit % init[i] == 0) {
			limit /= init[i];
		}
		if (limit == 1) {
			printf("%u\n", init[i]);
			return 0;
		}
	}

	/* Factors >= 7. */
	uint64_t step = 2;
	for (i = 7; i * i <= limit; i += step) {
		while (limit % i == 0) {
			limit /= i;
		}
		step = 6 - step;
	}

	if (limit == 1) {
		printf("%lu\n", i - step);
	} else {
		printf("%lu\n", limit);
	}
	return 0;
}
