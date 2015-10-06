/* There is a constant solution thanks to the expression of the sum of the
* integers and the sum of the squares:

    result(n) = n*(n+1)*(3*n^2-n-1)/6 - [n(n+1)/2]^2
              = n*(n+1)*(3*n^2-n-2)/12
              = [n*(n+1)/2]*[(2*n^2 + n^2-n - 2)/6]
              = [n*(n+1)/2]*[(n^2 + n*(n-1) - 1)/3]

The last expression has the advantage of keeping values low.
*/

#include <stdio.h>

int main() {
	int limit = 100;
	int result = limit * (limit+1) / 2;
	result *= (limit * limit + (limit * (limit-1) / 2) -1) / 3;

	printf("%d\n", result);
	return 0;
}
