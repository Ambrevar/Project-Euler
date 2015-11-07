/* The clever method is to add the highest value among the two children to every
element starting from the base of the triangle. The result will be in the root /
summit.

We use a triangular data structure to store the values, this allows for dynamic
input size as opposed to a regular matrix.
*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main() {
	/* Maximum size of the triangle. */
	int limit = 100;

	FILE *file = fopen("input.txt", "rb");
	if (!file) {
		return 1;
	}

	long *triangle = calloc(limit * (limit + 1) / 2, sizeof (long));
	if (!triangle) {
		fclose(file);
		return 1;
	}

	/* Fill the triangle from the file. */
	int i;
	long value;
	for (i = 0; fscanf(file, "%ld", &value) != EOF; i++) {
		triangle[i] = value;
	}

	/* The real size of the triangle is the solution to i = limit(limit+1)/2. */
	limit = (-1 + sqrt(1 + 8 * i)) / 2;
	/* i-th row starts at i*(i-1)/2. */
	i = 0;
	int row;
	int offset;
	int offset_next;
	for (row = limit - 2; row >= 0; row--) {
		offset = row * (row + 1) / 2;
		offset_next = offset + row + 1;
		for (i = 0; i <= row; i++) {
			if (triangle[offset_next + i] > triangle[offset_next + i + 1]) {
				/* Left child is higher than right child. */
				triangle[offset + i] += triangle[offset_next + i];

			} else {
				triangle[offset + i] += triangle[offset_next + i + 1];
			}
		}
	}

	printf("%ld\n", triangle[0]);

	fclose(file);
	free(triangle);
	return 0;
}
