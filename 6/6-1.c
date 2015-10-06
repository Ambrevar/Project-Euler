/* Large number algorithm. We develop the square of the sum into two loops. We
skip the square numbers because of the difference.

This is much slower than the constant solution

    f(n)=n*(n+1)*(3*n^2-n-2)/12

but may reach higher values.
*/

#include <stdio.h>

int main() {
	int limit = 100;
	int result = 0;

	int i, j;
	for (i = 1; i < limit; ++i) {
		for (j = i + 1; j <= limit; ++j) {
			result += i * j;
		}
	}

	result *= 2;

	printf("%d\n", result);
	return 0;
}
