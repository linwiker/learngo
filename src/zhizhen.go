package main

import "fmt"

//

//func main() {
//	str := "Go is a beautiful language!"
//	for pos, char := range str {
//		fmt.Printf("%d %c \n", pos, char)
//	}
//}

func main() {
	for i := 0; i<7 ; i++ {
		if i%2 == 0 { continue }
		fmt.Println("Odd:", i)
	}
}