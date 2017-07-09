package main

import "fmt"

const (
	FIZZ      = 3
	BUZZ 	   = 5
	FIZZBUZZ = 15
)

func main() {
	for i:= 0; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FIZZBUZZ")
		case i%FIZZ == 0:
			fmt.Println("FIZZ")
		case i%BUZZ == 0:
			fmt.Println("BUZZ")
		default:
			fmt.Println(i)
		}
	}
}
