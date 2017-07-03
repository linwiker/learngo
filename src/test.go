package main

import (
	"fmt"
)
var v int = 2
func main()  {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	fmt.Println(v)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
