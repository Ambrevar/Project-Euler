/* We use different algorithms here:

- We set a boolean table of indices which are sums of abundant numbers. It is
  sort of a sieve.

- We compute the sum of divisors s(n) using that when p does not divide a,
  s(a * p) = s(a) * s(p).
  When p is prime, s(p^n) = 1 + p + ... + p^n (geometrical sum).
*/

#include <stdio.h>
#include <stdbool.h>
#include <string.h>

#define ABUND_MAX 28123

/* Note that divisor_sum(1) == 0. */
unsigned int divisor_sum(unsigned int n) {
	unsigned int input = n;
	unsigned int sum = 1;
	unsigned int i;

	// We loop starting from 7 with a step of 4,2,4,2... But we need to treat the
	// frist primes as a special case.
	unsigned int init[3] = {2, 3, 5};
	for (i = 0; i < 3; i++) {
		unsigned int p = 1;
		while (n % init[i] == 0) {
			p = p * init[i] + 1;
			n /= init[i];
		}
		sum *= p;
	}

	unsigned int step = 2;
	for (i = 7; i * i <= n; i += step) {
		// `p` holds the divisor sum of i^a, where `a` is the highest power so that
		// i^a divides n.
		unsigned int p = 1;
		while (n % i == 0) {
			// Geometrical sum.
			p = p * i + 1;
			n /= i;
		}
		sum *= p;
		step = 6-step;
	}
	// Since we early-out the previous loop when we reach i*i > n, there are 2
	// possible outcomes: either n==1 or i<=n<i*i and n is prime. In the latter
	// case we need to add it to the result.
	if (n > 1) {
		sum *= 1 + n;
	}

	// Our definition of divisor sum of n does not include n itself.
	return sum - input;
}

int main() {
	unsigned int abundant[ABUND_MAX];
	unsigned int abundant_index = 0;
	memset(abundant, false, ABUND_MAX);

	bool abundant_sums[ABUND_MAX + 1];
	memset(abundant_sums, false, ABUND_MAX + 1);

	// Fill abundant table.
	unsigned int i, j;
	for (i = 12; i < ABUND_MAX; i++) {
		if (divisor_sum(i) > i) {
			abundant[abundant_index] = i;
			abundant_index++;
		}
	}

	for (i = 0; i < abundant_index; ++i) {
		unsigned long sum  = abundant[i] + abundant[0];
		for (j = 0; j <= i && sum <= ABUND_MAX; j++, sum = abundant[i] + abundant[j]) {
			abundant_sums[sum] = true;
		}
	}

	unsigned long result = 0;
	for (i = 1; i <= ABUND_MAX; ++i) {
		if (!abundant_sums[i]) {
			result += i;
		}
	}

	printf("%lu\n", result);
	return 0;
}
