/* Bruteforce version.
Remarkable idea with the 1/7 cyclic number:
  https://en.wikipedia.org/wiki/142857_(number)
*/

#include <stdbool.h>
#include <stdio.h>
#include <string.h>

bool comp(int n1, int n2) {
	char buf1[10];
	char buf2[10];
	memset(buf1, 0, 10);
	memset(buf2, 0, 10);

	while (n1 > 0) {
		buf1[n1 % 10]++;
		n1 /= 10;
	}
	while (n2 > 0) {
		buf2[n2 % 10]++;
		n2 /= 10;
	}
	int i;
	for (i = 0; i < 10; i++) {
		if (buf1[i] != buf2[i]) {
			return false;
		}
	}
	return true;
}

int main() {
	int limit = 12345;

	for (;; limit++) {
		int limit2 = 2*limit;
		if (comp(limit2, 3 * limit) &&
			comp(limit2, 4 * limit) &&
			comp(limit2, 5 * limit) &&
			comp(limit2, 6 * limit)) {
			break;
		}
	}

	printf("%d\n", limit);
	return 0;
}
