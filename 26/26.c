/* The cycle induced by 1/d cannot be longer than d-1. Thus we can go down from
the limit and if a `d` leads to a cycle of length `d-1`, we know it is the max
we are looking for. We can use this property to early-out.

We cannot filter out composite numbers since we have no garantee that no composite number between the last prime and the limit that has no bigger cycle.

See <http://en.wikipedia.org/wiki/Repeating_decimal>.
*/

#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define RANGE 1000

/* Return the cycle length of divid/quotient. */
unsigned int div_cycle_len(unsigned int divid, unsigned int quotient) {
	unsigned int i;
	unsigned int cycles[RANGE];
	memset(cycles, 0, RANGE * sizeof (unsigned int));
	unsigned int cycle_length = 0;

	/* WARNING: We do not check i for overflow, we assume we work within
	UINT_MAX */
	for (i = 1; divid != 0; i++) {

		if (divid >= quotient) {
			divid %= quotient;
		}

		/* If cycles[divid] == 0, `digit` cycle is not stored yet. */
		if (cycles[divid] == 0) {
			cycles[divid] = i;
		} else {
			/* Last time we hit `divid` was at position `cycles[divid]`. The cycle is
			current position minus last position. */
			cycle_length = i - cycles[divid];
			break;
		}

		divid *= 10;
	}

	return cycle_length;
}

int main() {
	unsigned int result = 0;
	unsigned int result_length = 0;

	unsigned int i;
	for (i = RANGE - 1; i > 0; i--) {
		unsigned int temp = div_cycle_len(1, i);
		if (temp > result_length) {
			result = i;
			if (temp == i - 1) {
				break;
			}
			result_length = temp;
		}
	}

	printf("%d\n", result);
	return 0;
}
