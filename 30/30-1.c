/* First we need to compute the upper bound: the power sum grows slower than the
number, so all we need is to find the intersection:

  n*9^power == 10^n

It is enough to find the number of digits n in the smallest number made of '9's
that is greater than its power sum. A good upper bound is then `n*9^power`.

An important speedup is to precompute all digit powers since they will be used
many times.

This version is simple but suboptimal.
*/
#include <stdio.h>

int decomp(int number, int *digit_powers) {
	int result = 0;

	while (number != 0) {
		result += digit_powers[number % 10];
		number /= 10;
	}

	return result;
}

int main() {
	int power = 5;
	int digit_powers[10];
	int i;
	for (i = 0; i < 10; i++) {
		int j;
		digit_powers[i] = i;
		for (j = 1; j < power; j++) {
			digit_powers[i] *= i;
		}
	}

	/* `upper_bound` is the max possible value for the given 'power'. */
	int upper_bound = digit_powers[9];
	int buffer = 9;
	for (i = 2; buffer < upper_bound; i++) {
		buffer = 10 * buffer + 9;
		upper_bound = i * digit_powers[9];
	}

	int result = 0;
	for (i = 10; i < upper_bound; ++i) {
		buffer = decomp(i, digit_powers);

		if (i == buffer) {
			result += i;
		}
	}

	printf("%d\n", result);
	return 0;
}
