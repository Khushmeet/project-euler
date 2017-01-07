package main

/*
Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

import "fmt"

func divByN(n, target int64) int64  {
    var p int64 = target/n
    return n*(p*(p+1))/2
}

func main() {
	var target int64 = 1000

    output := divByN(3, target) + divByN(5, target) - divByN(15, target)

	fmt.Println("Sum is :", output)
}
