package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// https://www.runoob.com/go/go-structures.html -- Go 语言结构体

func constructStruct() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

// 不过有时候没分隔符，看着不是很方便
// {Go 语言 www.runoob.com Go 语言教程 6495407}
// {Go 语言 www.runoob.com Go 语言教程 6495407}
// {Go 语言 www.runoob.com  0}

// 访问结构体成员
func accessMemberOfStruct() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	/* 打印 Book1 信息 */
	printStruct(Book1)
	printStruct(Book2)

	/* 打印 Book1 信息 */
	printBookPtr(&Book1)
	/* 打印 Book2 信息 */
	printBookPtr(&Book2)
}

// Book 1 title : Go 语言
// Book 1 author : www.runoob.com
// Book 1 subject : Go 语言教程
// Book 1 book_id : 6495407
// Book 2 title : Python 教程
// Book 2 author : www.runoob.com
// Book 2 subject : Python 语言教程
// Book 2 book_id : 6495700

// 结构体作为函数参数
func printStruct(book Books) {
	fmt.Println(book)
	fmt.Printf("%v \n", book)
}

// {Go 语言 www.runoob.com Go 语言教程 6495407}
// {Go 语言 www.runoob.com Go 语言教程 6495407}
// {Python 教程 www.runoob.com Python 语言教程 6495700}
// {Python 教程 www.runoob.com Python 语言教程 6495700}

// 结构体指针
func printBookPtr(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

// Book title : Go 语言
// Book author : www.runoob.com
// Book subject : Go 语言教程
// Book book_id : 6495407
// Book title : Python 教程
// Book author : www.runoob.com
// Book subject : Python 语言教程
// Book book_id : 6495700
