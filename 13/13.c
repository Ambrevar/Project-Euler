/* On the forum, it is told that only the first 11 sums are required. But what
about the carry? It is true in this particular case, but not in general. */

#include <stdio.h>
#include <stdlib.h>

int main() {
	unsigned int digits = 50;
	unsigned int number_count = 100;

	FILE *file = fopen("input.txt", "rb");
	if (file == NULL) {
		return 1;
	}

	char *table = NULL;
	unsigned int table_size = digits * number_count;
	table = malloc(table_size * sizeof (char));
	if (!table) {
		fclose (file);
		return 1;
	}
	fread(table, table_size, sizeof (char), file);
	fclose(file);

	/* Change char to integer value. */
	unsigned int i;
	for (i = 0; i < table_size; ++i) {
		table[i] -= '0';
	}

	char *result = NULL;
	result = malloc((digits + 2) * sizeof (char));

	/* Manual addition */
	unsigned int j, carry, sum;
	carry = 0;
	for (i = 1; i <= digits; ++i) {
		sum = 0;
		for (j = 0; j < number_count; ++j) {
			sum += table[j * digits + (digits - i)];
		}
		sum += carry;

		carry = sum / 10;
		result[52 - i] = sum % 10;
	}

	/* Add the last numbers. */
	sum /= 10;
	result[1] = sum % 10;
	sum /= 10;
	result[0] = sum % 10;

	unsigned long value = 0;
	for (i = 0; i < 10; i++) {
		value = value * 10 + result[i];
	}

	printf("%lu\n", value);

	free(result);
	free(table);
	return 0;
}
