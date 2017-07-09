package main

import "fmt"


func main() {
	//第一种方法
	for i := 1; i <= 6; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// 第二种方法
	var g string = "G"
	for i := 0; i < 6; i++ {
		fmt.Println(g)
		g += "G"
	}
}
