/* Naive algorithm. This is slower than the sliding version, but also simpler.

Trick: when we meet a '0', we can move the start index after the '0' since all
products found in between will be null.
*/

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main() {
	unsigned int size;
	unsigned int limit = 13;
	char *buffer;
	unsigned long long result;

	/* Store file to a buffer. */
	FILE *file = fopen("input.txt", "rb");
	if (file == NULL) {
		return 1;
	}
	fseek(file, 0, SEEK_END);
	size = ftell(file) - 1;
	fseek(file, 0, SEEK_SET);
	buffer = malloc(size * sizeof (char));
	fread(buffer, size, sizeof (char), file);
	fclose(file);

	/* Get integer value from chars. */
	unsigned int i;
	for (i = 0; i < size; ++i) {
		buffer[i] -= '0';
	}

	/* Store largest product. */
	unsigned int j;
	unsigned long long product;
	result = 0;
	unsigned int size_limit = size - limit;
	for (i = 0; i < size_limit; ++i) {

		product = 1;
		unsigned int ilimit = i + limit;
		for (j = i; j < ilimit; ++j) {
			if (buffer[j] == 0) {
				i = j;
				break;
			}
			product *= buffer[j];
		}

		if (product > result) {
			result = product;
		}
	}

	printf("%llu\n", result);

	free(buffer);
	return 0;
}
