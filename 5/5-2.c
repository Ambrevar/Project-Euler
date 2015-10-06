/* Using alternating sieve is a little faster. */
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>

int main() {
	unsigned int limit = 20;

	/* Result can be big; however it is bound by 'limit!' which is < 2^64 for
	limit==20. */
	uint64_t result = 1;

	/* See problem 3 for the sieve. */
	bool *sieve = calloc(limit + 1, sizeof (bool));

	unsigned int i, j, step;
	step = 4;

	for (j = 2; j <= limit; j *= 2) {
		/* Do nothing */
	}
	result *= j / 2;

	for (j = 3; j <= limit; j *= 3) {
		/* Do nothing */
	}
	result *= j / 3;

	for (i = 5; i <= limit; i += step) {
		step = 6 - step;
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
