/* This version uses the alternating sieve. */

#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <math.h>

/* The alternating sieve uses an alternating loop step of 2-4. Since primes > 3
are not multiple of 2 nor 3, they are dispatched every 2 and 4 integers.
Considering primes > 5 would require an unefficient function to compute the step.
However, the loop for prime multiples starts at i*2, not i*i.

This is slightly faster than the traditional sieve. However, it does not yield a
sieve table, only the prime table.
*/
uint32_t make_sieve_alt(uint32_t limit, uint32_t *primes) {
	uint32_t bound;
	if (limit > 1) {
		bound = limit * limit;
	} else if (limit == 1) {
		bound = 2;
	} else {
		bound = 1;
	}

	bool *sieve = calloc (bound + 1, sizeof (bool));
	if (sieve == NULL)
		return 0;
	
	sieve[0] = true;
	if (limit == 0) {
		return 0;
	}
	sieve[1] = true;
	primes[0] = 1;
	primes[1] = 2;
	primes[2] = 3;
	uint32_t count = 1;
	uint32_t i;
	uint32_t step = 4;
	for (i = 5; i <= limit; i+= step) {
		step = 6 - step;
		if (!sieve[i]) {
			primes[count++] = i;

			uint32_t j;
			for (j = i * 2; j < bound; j += i) {
				sieve[j] = true;
			}
		}
	}

	/* If we would only care about the first limit-th values, this loop would
	not be necessary. */
	for (i = limit + 1; i <= bound; i+=step) {
		step = 6 - step;
		if (!sieve[i]) {
			primes[count++] = i;
		}
	}

	free (sieve);
	return count;
}

int main() {
	uint64_t limit = 600851475143;

	uint32_t size = (uint32_t)round(sqrt ((double)limit));
	uint32_t gen_limit = (uint32_t)round(sqrt ((double)size));

	uint32_t *primes = malloc (size * sizeof (uint32_t));
	if (primes == NULL)
		return 1;
	
	uint32_t count = make_sieve_alt (gen_limit, primes);
	uint32_t i;

	for (i = count - 1; i > 1; i--)	{
		if (limit % primes[i] == 0) {
			printf ("%u\n", primes[i]);
			return 0;
		}
	}

	return 0;
}
