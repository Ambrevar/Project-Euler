/* We generate decimal palindromic numbers and check if they are binary
palindromic as well.
*/
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

bool bin_palindromic(unsigned int n) {
	unsigned int temp = n;

	unsigned int mask_high = 1;
	unsigned int mask_low = 1;

	/* Set mask_high */
	while (temp != 0) {
		temp /= 2;
		mask_high <<= 1;
	}
	mask_high >>= 1;

	/* Check if palindromic */
	while (mask_high >= mask_low) {
		/* Caution here: must compare to previous value. */
		if (((n & mask_high) == mask_high) != ((n & mask_low) == mask_low)) {
			return false;
		}

		mask_high >>= 1;
		mask_low <<= 1;
	}

	return true;
}

int main() {
	/* Sum of all first bin-palindromes below 10. */
	unsigned int result = 25;

	/* No need to check even numbers */
	unsigned int i;
	for (i = 1; i <= 999; i += 2) {
		unsigned int reverse_i, order = 1;
		order = 1;
		reverse_i = 0;

		{
			unsigned int temp = i;
			while (temp != 0) {
				reverse_i += temp % 10;
				reverse_i *= 10;
				order *= 10;
				temp /= 10;
			}
		}
		reverse_i /= 10;

		/* Insert 0 in the middle, i.e. 77, then 7007, then 700007, etc. */
		for (; reverse_i * order < 1000000; order *= 100) {
			unsigned int palindrome = i + reverse_i * order;
			if (bin_palindromic(palindrome)) {
				result += palindrome;
			}

			/* Odd digit number: insert 'j' in the middle. */
			if (i < 100) {
				unsigned int j;
				for (j = 0; j <= 9; j++) {
					palindrome = i + j * order + reverse_i * order * 10;
					if (bin_palindromic(palindrome)) {
						result += palindrome;
					}
				}
			}
		}
	}

	printf("%u\n", result);
	return 0;
}
