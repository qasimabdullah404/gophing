/*
TIMESTAMP: 20240606023207
PROBLEM STATEMENT:
	If we list all the natural numbers below 10 that are multiples of 3 or 5,
	we get 3, 5, 6 and 9. The sum of these multiples is 23.

	Find the sum of all the multiples of 3 or 5 below 1000.
*/

/*
SOLUTION: 233168
*/

package main

import "fmt"

func main() {
	var numbers []int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			numbers = append(numbers, i)
		}
	}

	fmt.Println(sum(numbers))
}

func sum(n []int) int {
	var sum int
	for _, number := range n {
		sum = sum + number
	}

	return sum
}
