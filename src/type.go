package main

import "fmt"

func main() {
    k, b := 7,8
    fmt.Printf("before k=%d,b=%d\n", k, b)
    if k :=1; k != -1 {
         b := 3
         fmt.Printf("inner k=%d,b=%d\n", k, b)
    }
    fmt.Printf("after k=%d,b=%d\n", k, b)
}
