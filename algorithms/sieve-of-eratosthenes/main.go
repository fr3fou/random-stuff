package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", findPrimes(n))
}

func findPrimes(n int) []int {
	// Create a list of consecutive integers from 2 through n
	// (2, 3, 4, ..., n)
	nums := make([]int, n, n)
	primes := make([]int, 0)

	for i := 2; i < n; i++ {
		nums[i] = i
	}

	// Initially, let p equal 2, the smallest prime number
	p := 2
	primes = append(primes, p)

	for {
		np := primeSieve(p, p, n, nums)

		// If np hasn't changed, our work is done
		if np == p {
			break
		}

		p = np
		primes = append(primes, p)
	}

	primes = primes[:len(primes)-2]

	return primes
}

// primeSieve takes in p (current num)
// i (incrementer / original p)
// n (maximum number)
// nums (list of the numbers)
func primeSieve(p, i, n int, nums []int) int {
	// Enumerate the multiples of p by counting in increments of p from 2p to n,
	// and mark them in the list (these will be 2p, 3p, 4p, ...;
	// the p itself should not be marked).

	// If we have reached the maximum,
	// return the first number greater than p that is not marked
	if p+i >= n {
		for _, v := range nums {
			if v > i && v != -1 {
				return v
			}
		}

		return p
	}

	// Mark in the list
	nums[p] = -1

	// Continue recursively
	return primeSieve(p+i, i, n, nums)
}
