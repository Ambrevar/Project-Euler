/* This is a combinatorial problem. The key is to notice that whatever route you
take, there will be the same number of 'downs' and 'rights', that is 20 and 20.

The wanted result is the combination of 20 in 40, i.e. \binom{40}{20}.

Pitfall: numbers are too big for 64-bits integers. We need high precision.
*/

#include <stdint.h>
#include <stdio.h>

int main() {
	int limit = 20;
	int limit_x2 = limit * 2;
	int i = 1;

	/* Compute factorials. */
	double f = 1;
	for (i = 1; i <= limit; i++) {
		f *= i;
	}
	double f2 = 1;
	for (i = limit+1; i <= limit_x2; i++) {
		f2 *= i;
	}

	printf("%.15g\n", f2/f);
	return 0;
}
