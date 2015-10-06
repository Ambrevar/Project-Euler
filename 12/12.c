/* The idea is to decompose into primes numbers (e.g. n = p1^a1 * p2^a2...).
Then the divisor count (divcount) is (a1+1)*(a2+1)...

Since it is a triangle number, i.e. an integer sum, it can be written n/2*(n+1).

If n is even: since n and n+1 have no divisors in common, the number or
divisor of n/2 * (n+1) is divcount(n/2) * divcount(n+1).

If n+1 is even: vice-versa.

We use some memoization to avoid redundant computing of divcount().
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

/* See problem 3. */
uint32_t make_sieve_alt(uint32_t limit, uint32_t *primes) {
	uint32_t bound;
	if (limit > 1) {
		bound = limit * limit;
	} else if (limit == 1) {
		bound = 2;
	} else {
		bound = 1;
	}

	bool *sieve = calloc(bound + 1, sizeof (bool));
	if (sieve == NULL) {
		return 0;
	}

	sieve[0] = true;
	if (limit == 0) {
		return 0;
	}
	sieve[1] = true;
	primes[0] = 1;
	primes[1] = 2;
	primes[2] = 3;
	uint32_t count = 3;
	uint32_t i;
	uint32_t step = 4;
	for (i = 5; i <= limit; i += step) {
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
	for (i = limit + 1; i <= bound; i += step) {
		step = 6 - step;
		if (!sieve[i]) {
			primes[count++] = i;
		}
	}

	free(sieve);
	return count;
}


unsigned int divcount(uint64_t n, uint32_t *primes, uint32_t count) {
	uint32_t i;
	uint32_t result = 1;
	uint32_t prod;
	for (i = 1; n > 1 && i < count; ++i) {
		prod = 1;
		while (n % primes[i] == 0) {
			n /= primes[i];
			prod++;
		}
		result *= prod;
	}

	return result;
}

int main() {
	unsigned int limit = 500;

	/* Let n be such that triangle = n*(n+1)/2. Then n < sqrt(triangle). Since
	triangle is assumed to be < 2^32, we have that n < 2^16. */
	uint32_t memoiz_size = 65536;
	uint32_t *memoization = calloc(memoiz_size, sizeof (uint32_t));
	if (!memoization) {
		return 1;
	}

	/* To lookup prime divisors of n, we only need to go up to sqrt(n) < sqrt
	(2^16) == 2^8. */
	/* TODO: setting size = 128 slows down primes access. Why? */
	uint32_t size = 256;
	uint32_t *primes = malloc(size * size * sizeof (uint32_t));
	if (!primes) {
		return 1;
	}
	uint32_t count = make_sieve_alt(size, primes);

	uint64_t i;
	uint64_t triangle = 0;
	for (i = 1;; i++) {
		triangle += i;

		uint64_t n = i;
		uint64_t n1 = n + 1;
		if (n % 2 == 0) {
			n /= 2;
		} else {
			n1 /= 2;
		}
		if (memoization[n] == 0) {
			memoization[n] = divcount(n, primes, count);
		}
		if (memoization[n1] == 0) {
			memoization[n1] = divcount(n1, primes, count);
		}
		uint64_t dcn = memoization[n];
		uint64_t dcn1 = memoization[n1];

		if (dcn * dcn1 > limit) {
			break;
		}

	}

	printf("%lu\n", triangle);

	free(primes);
	free(memoization);
	return 0;
}
