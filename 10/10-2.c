/* Alternative sieve with an alternating loop step of 2-4.
This can be faster in some circumstances than the classic sieve.

Note that the loop over prime multiples starts at 2*i and not i*i.
*/

#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <math.h>

/* Everytime we stop on a prime, we append it to the sum. */
uint64_t sum_sieve(uint32_t limit, uint32_t bound, bool *sieve) {
	uint64_t sum = 5;

	uint32_t i;
	unsigned int step;
	step = 4;
	for (i = 5; i <= limit; i += step) {
		step = 6 - step;

		if (!sieve[i]) {
			sum += i;
			uint32_t j;
			for (j = 2 * i; j <= bound; j += i) {
				sieve[j] = true;
			}
		}
	}

	/* If we would only care about the first limit-th values, this loop would
	not be necessary. */
	for (i = limit + 1; i <= bound; i += step) {
		step = 6 - step;
		if (!sieve[i]) {
			sum += i;
		}
	}

	return sum;
}

int main() {
	uint64_t limit = 2000000;
	uint64_t limit_sqrt = round(sqrt(limit));

	bool *sieve = calloc(limit + 1, sizeof (bool));
	if (sieve == NULL) {
		return 1;
	}

	uint64_t result = sum_sieve(limit_sqrt, limit, sieve);

	printf("%lu\n", result);
	return 0;
}
