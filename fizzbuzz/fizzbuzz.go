// The fizzbuzz program prints a list of integers up to a given length
// (default 20), replacing the number with fizz when it is divisible by 3,
// buzz when divisible by 5 and fizzbuzz when divisible by 3 and 5.
package main

import (
	"flag"
	"fmt"
	"strconv"
)

var limit = flag.Int("limit", 20, "size of the list to fizzbuzz")

func main() {
	flag.Parse()
	fizzbuzz := make([]string, 0, *limit)

	for i := 0; i < *limit; i++ {
		var x string
		switch {
		case i == 0:
			x = "0"
		case i%3 == 0 && i%5 == 0:
			x = "fizzbuzz"
		case i%3 == 0:
			x = "fizz"
		case i%5 == 0:
			x = "buzz"
		default:
			x = strconv.Itoa(i)
		}
		fizzbuzz = append(fizzbuzz, x)

	}

	fmt.Println(fizzbuzz)
}
