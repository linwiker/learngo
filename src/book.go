package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "linwiker"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 653210

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "linwiker"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 123012031

	/* 打印Book1信息 */
//	fmt.Println(Book1.title, Book1.author, Book1.subject, Book1.book_id)
	//
	//	//打印Book2信息
	//	fmt.Println(Book2.title, Book2.author, Book2.subject, Book2.book_id)
	printBook(&Book1)
	printBook(&Book2)

	s := []int {1, 2, 3}
	fmt.Println(s[1])
	z := make([]int, 3, 5)
	fmt.Println(len(z), cap(z), z[1])
}


func printBook( book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}