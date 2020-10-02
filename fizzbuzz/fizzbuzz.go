package main

import "fmt"

func main() {
	for i := 0; i <= 50; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(i int) {
	if i % 3 == 0 {
		fmt.Printf("Fizz")
	}
	if i % 5 == 0 {
		fmt.Printf("Buzz")
	}
	if i%5!=0 && i%3!=0	{
		fmt.Printf("number=%d", i)
	}
	fmt.Printf("\n")
}

