/* We could use a table of amicables which have already been processed. However
the frequence of amicable is not high enough for it to be worth it, performance
is typically lower with record.
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

unsigned int divisor_sum(unsigned int n) {
	unsigned int i;
	/* '1' is included in the result. */
	unsigned int result = 1;

	for (i = 2; i * i < n; ++i) {
		if (n % i == 0) {
			result += i + n / i;
		}
	}

	if (i * i == n) {
		result += i;
	}

	return result;
}


int main() {
	/* limit >= 2 */
	unsigned int limit = 100000;
	unsigned int result = 0;
	unsigned int n;
	unsigned int amicable;

	for (n = 2; n <= limit; n++) {
		amicable = divisor_sum(n);
		if (amicable > n && divisor_sum(amicable) == n) {
			result += n + amicable;
		}
	}

	printf("%u\n", result);
	return 0;
}
