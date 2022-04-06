package main

import "fmt"

func fizzbuzz(total int) {
	for nilai := 1; nilai <= total; nilai++ {
		if nilai%3 == 0 && nilai%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if nilai%3 == 0 {
			fmt.Println("Fizz")
		} else if nilai%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(nilai)
		}
	}
}

func main() {
	fizzbuzz(100)
}
