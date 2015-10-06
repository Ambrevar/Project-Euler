/* Simple sieve. Cf. problem 3.

The number of primes below limit = 2000000 is

  Pi(limit) = limit/ln(limit) ~= 137849

The sum is bound:

  sum < limit*(limit+1)/2 - Pi(limit)*(Pi(limit)+1)/2
      ~= 1 990 499 795 305

So we need an unsigned 64 bit integer to store the sum.
*/

#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <math.h>

/* Everytime we stop on a prime, we append it to the sum. */
uint64_t sum_sieve(uint32_t limit, uint32_t bound, bool *sieve) {
	uint64_t sum = 0;

	uint32_t i;
	for (i = 2; i <= limit; i++) {

		if (!sieve[i]) {
			sum += i;
			uint32_t j;
			for (j = i * i; j <= bound; j += i) {
				sieve[j] = true;
			}
		}
	}

	/* If we would only care about the first limit-th values, this loop would
	not be necessary. */
	for (i = limit + 1; i <= bound; i++) {
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
