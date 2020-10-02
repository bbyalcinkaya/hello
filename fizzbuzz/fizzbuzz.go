package main

import "fmt"

func main() {
	for i := 0; i <= 50; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(i int) {
	if i % 3 == 0 {
		fmt.Printf("Fizz\n")
	} else if i % 5 == 0 {
		fmt.Printf("Buzz\n")
	} else {
		fmt.Printf("number=%d\n", i)
	}
}
