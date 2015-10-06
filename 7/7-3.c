/* We stop the main loop at the square root + we use the alternating sieve.
This is even faster.
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main() {
	unsigned int limit = 10001;

	/* Size is the square root of the total_size of the table. */
	unsigned int size = 342;
	unsigned int total_size = size * size;
	unsigned int i;
	unsigned int j;
	unsigned int count = 2;

	/* Use calloc to make sure the table is initiated to zero-values. */
	bool *sieve = calloc(total_size, sizeof (bool));

	/* '0' and '1' are not prime. */
	sieve[0] = true;
	sieve[1] = true;
	unsigned int step = 4;
	for (i = 5; i < size && count < limit; i+=step) {
		step = 6-step;
		if (!sieve[i]) {
			count++;

			for (j = i * i; j < total_size; j += i) {
				sieve[j] = true;
			}
		}
	}

	for (; i < total_size && count < limit; i += step) {
		step = 6-step;
		if (!sieve[i]) {
			count++;
		}
	}

	free(sieve);
	printf("%d\n", i - step);
	return 0;
}
