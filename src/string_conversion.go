package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "88888"
	var an int
	var error error

	fmt.Println(strconv.IntSize)
	fmt.Println(strconv.Itoa(100))
	an, error  = strconv.Atoi(orig)
	fmt.Println(an, error)
}
