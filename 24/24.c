/* The number of permutations of a set of n elements is n!.

We seek the Nth permutation. n! can be divided into n segments of size (n-1)!.
First segment will be all permutation starting with the first digit, second
segment with second digit, and so on.

From the segment where N lies, we can infer the first digit in the result, then
work on a smaller set without this digit where we look for permutation `N -
(segment start)`.
*/

#include <stdio.h>

void table_delete(char *table, int *size, char element) {
	int i;
	for (i = 0; i < *size; ++i) {
		if (table[i] == element) {
			while (i < *size) {
				table[i] = table[i + 1];
				i++;
			}
		}
	}
	(*size)--;
}

int main() {
#define INITSIZE 10

	char dictionary[INITSIZE] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}; /* Must be ordered. */
	char result[INITSIZE] = {0};
	int size = INITSIZE;

	/* A segment is the number of permutation starting with a specific digit. */
	int segment_size = 1;
	int i;
	for (i = 2; i < INITSIZE; i++) {
		segment_size *= i;
	}

	/* First element is at index 0, it is more convenient for table indices and
	modulos. */
	int limit = 1000000;
	limit--;
	/* Permutation count in the dictionary is `size * segment_size`. */
	if (limit > size * segment_size) {
		printf("ERROR: limit too large.\n");
		return 1;
	}

	int division;
	for (i = 0; i < INITSIZE; ++i) {
		division = limit / segment_size;
		result[i] = dictionary[division];
		table_delete(dictionary, &size, dictionary[division]);

		if (division != 0) {
			limit %= segment_size;
		}

		if (size != 0) {
			segment_size /= size;
		}
	}

	for (i = 0; i < INITSIZE; i++) {
		printf("%d", result[i]);
	}
	puts("");
	return 0;
}
