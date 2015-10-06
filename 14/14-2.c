/* This one uses memoization in a descending loop. It is slower than the
ascending loop.
*/

#include <stdio.h>
#include <stdlib.h>

int main() {
	unsigned int limit = 1000000;

	unsigned int chain_length;

	unsigned int chain_max = 0;
	unsigned int index_max = 0;
	unsigned long collatz;

	unsigned int *memoiz = calloc(limit, sizeof (int));
	memoiz[1] = 1;

	unsigned int i;
	for (i = limit; i != 0; i--) {
		collatz = i;
		chain_length = 0;

		while (collatz > limit || memoiz[collatz] == 0) {
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
