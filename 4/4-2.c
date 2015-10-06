/* If the number of digits is known, it is possible the factorize the product so
that less value will be tested.

Let's consider a 6-digits palindrome P = a*b; we can write P='xyzzyx'.

P = x*100000 + y*10000 + z*1000 + z*100 + y*10 + x
  = x*100001 + y*10010 + z*1100
  = 11*(9091x + 910y + 100z)

So if a is not multiple of 11, then b must be. This way only multiple of 11 need
to be tested for one digit, and we loop on the other.
*/

#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

bool is_palindrome(int n) {
	int r = 0;
	if (n % 10 == 0) {
		return false;
	}

	while (n > r) {
		r = r * 10 + n % 10;
		n /= 10;
	}
	if (r == n || r == n * 10 + r % 10) {
		return true;
	}
	return false;
}

int main() {

	/* If you choose a different step, the start value changes as well. A step of
	1 will work for any palindrom. A step of 11 will work for 6-digits
	palindroms. */

	/* const int m_step = 1; */
	/* const int m_max = 999; */
	/* const int m_min = 100; */

	const int m_step = 11;
	const int m_max = 990;
	const int m_min = 110;

	const int x_min = 100;
	const int x_max = 999;
	int m = m_max;
	int x = m;
	int result = 0;

	while (m >= m_min && x >= m_min && is_palindrome(m * x) == false) {
		if (x > x_min) {
			x--;
		} else {
			m -= m_step;
			x = 999;
		}
	}

	/* Once a palindrom has been found, we must check if there is another possible
	greater palindrom with a lower m. */
	if (m >= m_min && x > x_min) {
		result = m * x;

		m -= m_step;
		x = x_max;
		int prod;
		while (m * x_max > result) {
			prod = m * x;
			if (prod > result) {
				if (is_palindrome(prod) == true) {
					result = prod;
					m -= m_step;
					x = 999;
				} else {
					x--;
				}
			} else {
				m -= m_step;
				x = x_max;
			}
		}
	}

	printf("%d\n", result);

	return 0;
}
