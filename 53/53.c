/* We use the pascal triangle property to move to the next row:

  (n r) = (n-1 r) + (n-1 r-1)

We traverse the triangle by sticking to the lowest value above the limit. This
ensures that we don't work with big numbers, since the center value is the
highest.

Since every row is symmetric, with r-th value of the row being the first one
above >= limit, the number of values >= limit one the row is

  n+1 - 2*r
*/

#include <stdbool.h>
#include <stdio.h>

int main() {
	int limit = 1000000;

	int n;
	int r = 1;
	int center = 2;
	int center_left = 1;

	/* Find the first central number >= limit. */
	for (n = 2; center <= limit && n <= 100; ) {
		center = center + center_left;
		n++;
		r = (n + 1) / 2;
		center_left = center * r / (n - r + 1);
	}

	int result = 0;
	for (; n <= 100; ) {
		/* Move 'center' to the lowest value >= limit. */
		while (center_left >= limit) {
			int buf = center_left;
			r--;
			center_left = center_left * r / (n - r + 1);
			center = buf;
		}

		/* There are 'n+1' numbers at row 'n'. */
		result += n + 1 - 2 * r;

		/* Go down. */
		center = center + center_left;
		n++;
		center_left = center * r / (n - r + 1);
	}

	printf("%d\n", result);
	return 0;
}
