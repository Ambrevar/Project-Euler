/* We use all the techniques of the simple version, while using digit
permutations instead of the whole number set to reduce the size of the traversed
set. Indeed, permutations have the same power sum.
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

/* This will traverse all integer digits within range that are unique
permutation-wise. The algorithm is simple: we increase digit values in numberic
order so that

  array[i] <= array[i+1].

Everytime we had 1 digit we reset all digits to 0.
*/
bool next_digitset(int *array, int size, int *digit_count) {
	int i = 0;
	while (i < size && array[i] == 9) {
		i++;
	}
	if (i == size) {
		// No more permutation.
		return false;
	}
	if (i >= *digit_count) {
		// New digit is '0' since array was calloc'ed.
		(*digit_count)++;
	} else {
		array[i] += 1;
	}
	i--;
	while (i >= 0) {
		array[i] = array[i + 1];
		i--;
	}
	return true;
}

/* WARNING: `array` integrity is not guaranteed. You should copy the array
before calling this function. */
bool find_number_in_array(unsigned long n, int *array, int size) {
	int digit;

	while (n > 0) {
		digit = n % 10;
		int i;
		int oldsize = size;
		for (i = 0; i < size; i++) {
			if (array[i] == digit) {
				array[i] = array[size - 1];
				size--;
				break;
			}
		}

		if (size == oldsize) {
			/* Element not found. */
			return false;
		}
		n /= 10;
	}

	if (size != 0) {
		/* Check if we matched all elements in the array. */
		return false;
	}
	return true;
}

int main() {
	int power = 5;
	unsigned long digit_powers[10];
	int i;
	for (i = 0; i < 10; i++) {
		int j;
		digit_powers[i] = i;
		for (j = 1; j < power; j++) {
			digit_powers[i] *= i;
		}
	}

	int digit_bound = 1;
	{
		/* `upper_bound` is the max possible value for the given 'power'.
		WARNING: We assume this does not overflow.
		*/
		unsigned long buffer = 9;
		unsigned long upper_bound = digit_powers[9];
		for (i = 2; buffer < upper_bound; i++) {
			buffer = 10 * buffer + 9;
			upper_bound = i * digit_powers[9];
			digit_bound++;
		}
	}

	/* Digit order does not matter. We store from the beginning to simplify access. */
	int *array = calloc(digit_bound, sizeof (int));
	int *buffer = calloc(digit_bound, sizeof (int));
	array[0] = 0;
	int digit_count = 1;

	unsigned long result = 0;
	while (true) {
		unsigned long power_sum = 0;
		for (i = 0; i < digit_count; i++) {
			power_sum += digit_powers[array[i]];
		}

		/* We need to use a copy of the array since we want to keep it unchanged. */
		memcpy(buffer, array, digit_count * sizeof (int));
		if (find_number_in_array(power_sum, buffer, digit_count)) {
			result += power_sum;
		}

		if (next_digitset(array, digit_bound, &digit_count) == false) {
			// No more permutations.
			break;
		}
	}

	free(array);
	free(buffer);
	/* We remove '1' since it is not a sum. */
	printf("%lu\n", result - 1);
	return 0;
}
