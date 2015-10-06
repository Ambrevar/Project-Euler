/* Two ways of doing this: either implement everything by yourself, or use the
libc functions, like `qsort` and `sscanf`.

`qsort` is efficient, but `sscanf` is too heavy for our goal.
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static int cmpstringp(const void *p1, const void *p2) {
	/* The actual arguments to this function are "pointers to pointers to char",
	but strcmp(3) arguments are "pointers to char", hence the following cast plus
	dereference */
	return strcmp(*(char *const *)p1, *(char *const *)p2);
}


int main() {
	FILE *file = fopen("p022_names.txt", "rb");
	if (file == NULL) {
		printf("ERROR: could not open file!\n");
		return 1;
	}

	unsigned int file_size;
	fseek(file, 0, SEEK_END);
	file_size = ftell(file);
	fseek(file, 0, SEEK_SET);

	/* `buffer` is used to store file in memory. */
	char *buffer = malloc((file_size + 1) * sizeof (char));
	fread(buffer, file_size, sizeof (char), file);
	fclose(file);
	buffer[file_size] = '\0';

	/* We keep the parsing simple since we assume full control over the input */
	/* file. */
	char *parsed_string = strtok(buffer, "\",");

	/* Table for the strings. */
	#define NAME_COUNT 5163
	char **table = malloc(NAME_COUNT * sizeof (char *));
	/* Point to last string in table. */
	char **table_cell = table;

	int count = 0;
	while (parsed_string != NULL) {
		*table_cell = parsed_string;
		count++;
		parsed_string = strtok(NULL, "\",");
		table_cell++;
	}
	qsort(table, count, sizeof (char *), cmpstringp);

	unsigned long long result = 0;
	result = 0;
	int i;
	for (i = 0; i < count; ++i) {
		unsigned long long sum = 0;
		int j;
		for (j = 0; table[i][j] != '\0'; ++j) {
			sum += table[i][j] - 'A' + 1;
		}
		result += sum * (i + 1);
	}

	free(table);
	free(buffer);
	printf("%llu\n", result);
	return 0;
}
