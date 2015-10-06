/* Trick: use the direct expession of the integer sum:

    sum from 1 to N = N (N+1) / 2

Pitfall: multiples of 15 appear twice since 15 is both a multiple of 3 and 5.
Hence the subtraction.
*/
#include <stdio.h>

int main() {
	int limit = 999;
	int limit_3 = limit / 3;
	int limit_5 = limit / 5;
	int limit_15 = limit / 15;

	printf("%d\n", (3 * limit_3 * (limit_3 + 1) + 5 * limit_5 * (limit_5 + 1) - 15 * limit_15 * (limit_15 + 1)) / 2);
	return 0;
}
