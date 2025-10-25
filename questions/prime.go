package main

import (
	"fmt"
	"os"
	"strconv"
)

// IsPrime reports whether n is a prime number.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: prime <integer>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid integer:", os.Args[1])
		os.Exit(1)
	}

	if IsPrime(n) {
		fmt.Printf("%d is prime\n", n)
	} else {
		fmt.Printf("%d is not prime\n", n)
	}
}
