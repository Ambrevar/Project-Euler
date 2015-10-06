/* The only difference here is that we stop the main loop at the square root.
This is a tad faster.
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
	unsigned int count = 0;

	/* Use calloc to make sure the table is initiated to zero-values. */
	bool *sieve = calloc(total_size, sizeof (bool));

	/* '0' and '1' are not prime. */
	sieve[0] = true;
	sieve[1] = true;
	for (i = 2; i < size && count < limit; i++) {
		if (!sieve[i]) {
			count++;

			for (j = i * i; j < total_size; j += i) {
				sieve[j] = true;
			}
		}
	}

	i = size;
	if (i % 2 == 0) {
		i++;
	}
	for (; i < total_size && count < limit; i += 2) {
		if (!sieve[i]) {
			count++;
		}
	}

	printf("%d\n", i - 2);

	free(sieve);
	return 0;
}
