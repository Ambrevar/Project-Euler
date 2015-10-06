/* Trick: Use memoization. */

#include <stdio.h>
#include <stdlib.h>

int main() {
	unsigned int limit = 1000000;

	unsigned int chain_length;

	unsigned int chain_max = 0;
	unsigned int index_max = 0;
	unsigned long collatz;

	unsigned int *memoiz = malloc(limit * sizeof (int));
	/* Collatz(0) produces a chain of length 0. */
	memoiz[0] = 0;
	
	unsigned int i;
	for (i = 1; i <= limit; ++i) {
		collatz = i;
		chain_length = 0;

		while (collatz >= i && collatz != 1) {
			if (collatz % 2 == 0) {
				collatz /= 2;
			} else {
				collatz = 3 * collatz + 1;
			}
			chain_length++;
		}

		memoiz[i] = chain_length + memoiz[collatz];

		if (memoiz[i] > chain_max) {
			index_max = i;
			chain_max = memoiz[i];
		}
	}

	printf("%d\n", index_max);
	free(memoiz);
	return 0;
}
