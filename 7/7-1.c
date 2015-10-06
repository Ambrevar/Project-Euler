/* Just the sieve. Cf. problem 3. The sieve limit is given by a rough
approximation:

  Pi(n) ~= n/ln(n)

where Pi(n) is the approximate number of primes below n.
For Pi(n) = 10001, n = 116684.
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main() {
	unsigned int limit = 10001;

	unsigned int size = 116684;
	unsigned int i;
	unsigned int j;
	unsigned int count = 0;

	/* Use calloc to make sure the table is initiated to zero-values. */
	bool *sieve = calloc(size, sizeof (bool));

	/* '0' and '1' are not prime. */
	sieve[0] = true;
	sieve[1] = true;
	for (i = 2; i < size && count < limit; i++) {
		if (!sieve[i]) {
			count++;

			for (j = i * i; j < size; j += i) {
				sieve[j] = true;
			}
		}
	}

	printf("%d\n", i - 1);

	free(sieve);
	return 0;
}
