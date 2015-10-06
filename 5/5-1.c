/* See problem 3 for the sieve.

The result is the multiplication of every prime factor <= limit, each of them to
their respective highest exponent found in the decomposition of the numbers <=
20.

Algorithm: for every prime number p under the limit, we multiply the result by
the highest power of p below the limit.
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>

int main() {
	unsigned int limit = 20;
	/* Result can be big; however it is bound by 'limit!' which is < 2^64 for
	limit==20. */
	uint64_t result = 1;

	bool *sieve = calloc(limit + 1, sizeof (bool));

	unsigned int i, j;
	for (i = 2; i <= limit; i++) {
		if (!sieve[i]) {
			sieve[i] = true;

			for (j = i; j <= limit; j *= i) {
				/* Do nothing */
			}
			result *= j / i;

			for (j = i * i; j <= limit; j += i) {
				sieve[j] = true;
			}
		}
	}

	free(sieve);
	printf("%lu\n", result);
	return 0;
}
