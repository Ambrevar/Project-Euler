/* This is a combinatorial problem. The key is to notice that whatever route you
take, there will be the same number of 'downs' and 'rights', that is 20 and 20.

The wanted result is the combination of 20 in 40, i.e. \binom{40}{20}.

Pitfall: numbers are too big for 64-bits integers. We need big numbers.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var limit int64
	limit = 20
	b := big.NewInt(1)
	fmt.Println(b.Binomial(2*limit, limit))
}
