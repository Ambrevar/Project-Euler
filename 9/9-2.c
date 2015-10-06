/* a,b,c < 1000 so a*b*c < 10^9 == 1 billion, which an unsigned 32 bits integer
can hold.

The smart way (without programming):

From a^2 + b^2 = c^2, we invoke Euclid's formula with m > n:

    a = 2mn
    b = m^2 - n^2
    c = m^2 + n^2

Note that Euclid's formula does not specify whether a < b, which is not always
the case.

    a + b + c = 1000
    2mn + (m^2 - n^2) + (m^2 + n^2) = 1000
    2mn + 2m^2 = 1000
    2m(m+n) = 1000; m(m+n) = 500

Since m>n>0;

    m = 20
    n = 5

Hence the result:

    a = 200, b = 375, c = 425
*/

#include <stdio.h>
#include <stdint.h>

int main() {
	uint32_t limit = 1000;

	uint32_t m;
	uint32_t n;
	uint32_t half = limit / 2;
	uint32_t m2;
	uint32_t n2;

	for (n = 1;; n++) {
		/* Step by 2 since m and n have opposite parity. */
		for (m = n + 1;; m += 2) {
			uint32_t prod = m * (m + n);
			if (prod == half) {
				goto end;
			} else if (prod > half) {
				break;
			}
		}
	}
	end:

	m2 = m * m;
	n2 = n * n;
	printf("%d\n", 2 * m * n * (m2 - n2) * (m2 + n2));

	return 0;
}
