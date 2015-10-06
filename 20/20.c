/* We store the factorial result in a decimal table. A trick would be to change
computation from base 10 to base 2^N, with N>3. A higher base means fewer loops,
and a power of 2 means we can save on modulo and division operations. Sadly we
would need to convert this result to base 10, which requires big number support
(or another array). A higher base that is a power of 10 does not need big number
support to convert back to base 10, but then we lack the power of 2
optimization.
*/

#include <stdlib.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
	if (argc != 2) {
		fprintf(stderr, "Usage: %s <n>\n", argv[0]);
		return 0;
	}

	unsigned int limit = atoi(argv[1]);
	if (limit < 3) {
		fprintf(stderr, "Argument should be >= 3.\n");
		return 1;
	}

	/* `table` holds the decimal digits of the products. It cannot grow above */
	/* log_10(limit) + (limit-2)*log_10(limit) <= limit * limit. */
	unsigned int *table = malloc(limit * limit * sizeof (int));
	if (!table) {
		fprintf(stderr, "Allocation error.");
		return 2;
	}

	/* Init. */
	unsigned int i;
	unsigned int idx_bound;
	unsigned int init = limit;
	for (i = 0; init > 0; i++) {
		table[i] = init % 10;
		init /= 10;
		idx_bound = i;
	}

	/* Algorithm for 2 digits multipliers. */
	for (i = limit - 1; i >= 2; i--) {
		int carry = 0;

		unsigned int j;
		for (j = 0; j <= idx_bound || carry != 0; j++) {
			if (j > idx_bound) {
				idx_bound = j;
			}
			int product = (table[j] * i + carry);
			table[j] = product % 10;
			carry = product / 10;
		}
	}

	/* Compute the final sum. */
	unsigned long long sum = 0;
	for (i = 0; i <= idx_bound; i++) {
		sum += table[i];
	}

	printf("%llu\n", sum);
	free(table);
	return 0;
}
