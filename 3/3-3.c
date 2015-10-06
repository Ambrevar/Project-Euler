/* This version uses the sieve. See 3-2.

Algorithm: once the sieve has been generated, loop over all prime numbers
starting from the end. The first one to divide the input is the result.
*/

#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <math.h>

/*
The generation can stop at sqrt(len(sieve)), but computing sqrt() is expensive,
the we use the argument the other way around: it is the value of sqrt(), and we
generate a sieve of size limit*limit (max computed value is limit*limit-1).
'sieve' is false if index is prime, true otherwise. This may sound unintuitive,
but this is because Go's default value for arrays is false. 'primes' holds all
prime numbers in a non-sparse table. The n-th prime number is primes[n]. By
convention, primes[0] is 1. It returns the number of primes in primes.
*/
uint32_t make_sieve(uint32_t limit, bool *sieve, uint32_t *primes) {
	uint32_t bound;
	if (limit > 1) {
		bound = limit * limit;
	} else if (limit == 1) {
		bound = 2;
	} else {
		bound = 1;
	}

	sieve[0] = true;
	if (limit == 0) {
		return 0;
	}
	sieve[1] = true;
	primes[0] = 1;
	uint32_t count = 1;
	uint32_t i;
	for (i = 2; i <= limit; i++) {
		if (!sieve[i]) {
			primes[count++] = i;

			uint32_t j;
			for (j = i * i; j < bound; j += i) {
				sieve[j] = true;
			}
		}
	}

	/* If we would only care about the first limit-th values, this loop would
	not be necessary. */
	for (i = limit + 1; i < bound; i++) {
		if (!sieve[i]) {
			primes[count++] = i;
		}
	}

	return count;
}

int main() {
	uint64_t limit = 600851475143;

	uint32_t size = (uint32_t)round(sqrt((double)limit));
	uint32_t gen_limit = (uint32_t)round(sqrt((double)size));

	bool *sieve = calloc(size, sizeof (bool));
	if (sieve == NULL) {
		return 1;
	}
	uint32_t *primes = malloc(size * sizeof (uint32_t));
	if (primes == NULL) {
		free(sieve);
		return 1;
	}

	uint32_t count = make_sieve(gen_limit, sieve, primes);
	uint32_t i;

	for (i = count - 1; i > 1; i--) {
		if (limit % primes[i] == 0) {
			printf("%u\n", primes[i]);
			goto exit;
		}
	}

	exit:
	free(sieve);
	free(primes);
	return 0;
}
