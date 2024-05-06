package main

import (
	"fmt"
)

func isPrime(num int) bool {
	// чуть ускоренное решето Эратосфена
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func primesNum(min, max int) []int {
	primes := []int{}
	for i := min; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	primes := primesNum(1, 294)
	fmt.Println(primes)
}
