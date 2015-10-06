/* a,b,c < 1000 so a*b*c < 10^9, 1 billion, which an unsigned 32 bits integer
can hold.

This brute force solution is extremely slow compared to the smart way. It is
kept for the sole purpose of comparing performances.

Note the use of a 'goto' here: this is the only way to break multiple loops at
once in C. Another non-equivalent possibility would be to use a flag in all
loops, but that is more verbose and uses more CPU cycles. 'goto' have their use.
*/

#include <stdio.h>
#include <stdint.h>

int main() {
	uint32_t limit = 1000;

	uint32_t a;
	uint32_t b;
	uint32_t c;

	for (a = 0; a < limit; ++a) {
		for (b = a + 1; b < limit; ++b) {
			for (c = b + 1; c < limit; ++c) {
				if (a + b + c == limit && a * a + b * b == c * c) {
					goto end;
				}
			}
		}
	}
	end:

	printf("%d\n", a * b * c);

	return 0;
}
