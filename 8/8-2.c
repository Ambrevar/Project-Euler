/* We compute the full product only once, then we slide digit by digit: each
time we divide the product by the previously first digit and we multiply by the
new last digit.

This is faster than the non-sliding method.

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

	unsigned int size_limit = size - limit;

	/* Store largest product. */
	unsigned long long product = 0;
	for (i = 0; i < size_limit; ++i) {
		/* In case the new end digit is a 0, we skip it until product is non-null
		again. */
		unsigned int next_digit = i + limit;
		unsigned int j;
		while (product == 0 && i < size_limit) {
			/* If on previous loop the product was made null because of the last
			digit, we skip it. Using a 'while' instead of an 'if' is a free speed gain
			in case of several '0' spaced by 'limit' digits. */
			while (buffer[next_digit - 1] == 0 && i < size_limit - limit) {
				i += limit;
				next_digit = i + limit;
			}

			product = 1;
			for (j = i; j < next_digit; j++) {
				product *= buffer[j];
				/* If we step on a zero, we shift right after it and recompute the
				product. */
				if (buffer[j] == 0) {
					i = j + 1;
					next_digit = i + limit;
					break;
				}
			}
		}

		if (i >= size_limit) {
			break;
		}

		if (product > result) {
			result = product;
		}

		product *= buffer[next_digit];
		product /= buffer[i];
	}

	printf("%llu\n", result);

	free(buffer);
	return 0;
}
