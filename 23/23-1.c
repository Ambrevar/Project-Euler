/* First we find all abundant numbers. Then for all numbers i we sum those that
have no abundant numbers j below i/2 for which i-j is also abundant.
*/

#include <stdio.h>
#include <stdbool.h>
#include <string.h>

#define ABUND_MAX 28123

bool is_abundant(int number) {
	int result = 1;
	int i;

	for (i = 2; i * i < number; i++) {
		if (number % i == 0) {
			result += i + number / i;
		}
		/* Early out is too slow in this case.
		if (result > number) {
			return true;
		}
		*/
	}

	if (i * i == number) {
		result += i;
	}

	if (result > number) {
		return true;
	}

	return false;
}


int main() {
	unsigned long result = 0;
	bool tab_abundant[ABUND_MAX];
	memset(tab_abundant, false, ABUND_MAX);

	// Fill abundant table.
	int i, j;
	for (i = 12; i < ABUND_MAX; i++) {
		if (tab_abundant[i] != true && is_abundant(i) == true) {
			for (j = 1; i * j < ABUND_MAX; j++) {
				tab_abundant[j * i] = true;
			}
		}
	}

	for (i = 1; i < ABUND_MAX; i++) {
		// Try to find a `j` that would discard `i` from the result.
		for (j = 1;
			j <= i/2 &&
			(tab_abundant[j] == false ||
			tab_abundant[i - j] == false);
			j++) {
			;
		}

		if (j > i/2) {
			result += i;
		}
	}

	printf("%lu\n", result);
	return 0;
}
