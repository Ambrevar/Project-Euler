/* We perform the big number computation manually. We use big endian notation.

Interesting property: if we consider the recursive sum of the digits in 2^n,
then we get the sequence 1,2,4,8,7,5,1,2,4,8,...
However, it is no use for the following Euler problem.

Proof sketch:
2 * n = 2 * (d1*10^a1 + d2*10^a2...)
      = 2 * d1*10^a1 + 2 * d2*10^a2...

rds() means "recursive digit sum".

rds(2 * d*10^a) = 2 * rds(d)  (This is easy to show for every d in [0-9].)

Thus rds(2*n) = rds(2 * d1*10^a1 + 2 * d2*10^a2...)
              = 2*rds(d1*10^a1) + 2*rds(d2*10^a2) + ...
              = 2 * [rds(d1*10^a1) + rds(d2*10^a2) + ... ]
              = 2 * rds(d1*10^a1 + d2*10^a2...)
              = 2 * rds(n)

This boils down to proving it works for one-digit numbers. QED.
*/
#include <stdlib.h>
#include <stdio.h>



int main() {
	/*  Limit >= 1 */
	unsigned int limit = 1000;

	unsigned char *digits = malloc((limit + 1) / 3 * sizeof (char)); /* Sure of size ? */
	unsigned int i;
	unsigned int j;
	unsigned int max_index;
	unsigned int carry;
	unsigned int result;
	unsigned long long sum;

	digits[0] = 1;
	max_index = 0;

	for (i = 0; i < limit - 1; ++i) {
		carry = 0;

		for (j = 0; j <= max_index || carry != 0; ++j) {
			if (j > max_index) {
				max_index = j;
			}

			result = (digits[j] * 2 + carry);
			digits[j] = result % 10;

			carry = result / 10;
		}
	}

	/* Let's compute the final sum. */
	carry = 0;
	sum = 0;

	for (j = 0; j <= max_index || carry != 0; ++j) {
		if (j > max_index) {
			max_index = j;
		}
		result = (digits[j] * 2 + carry);
		digits[j] = result % 10;

		sum += digits[j];
		carry = result / 10;
	}

	printf("%llu\n", sum);
	free(digits);
	return 0;
}
