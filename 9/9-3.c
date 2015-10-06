/* a,b,c < 1000 so a*b*c < 10^9 == 1 billion, which an unsigned 32 bits integer
can hold.

The smart way again, where we compute m directly from n. If it yields an
integer, then we found value that fits.

We use only one loop. This is even faster.
*/

#include <stdio.h>
#include <stdint.h>
#include <math.h>

int main() {
	uint32_t limit = 1000;
	uint32_t limit_sqrt = sqrt(1000);

	uint32_t m;
	uint32_t n;
	uint32_t twice = limit * 2;
	uint32_t m2;
	uint32_t n2;

	for (n = 1; n < limit_sqrt; n++) {
		double ms = (sqrt(n * n + twice) - n) / 2;
		double ms_rounded = round(ms);
		if (ms == ms_rounded) {
			m = (uint32_t)round(ms);
			break;
		}
	}

	m2 = m * m;
	n2 = n * n;
	printf("%d\n", 2 * m * n * (m2 - n2) * (m2 + n2));

	return 0;
}
