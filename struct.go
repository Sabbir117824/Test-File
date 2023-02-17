package main

import "fmt"

type Book struct{//struct syntex
	Title string
	Author string
	ISBN string
	Price float32
	Page int
}
func main(){
	var b1 Book// struct varriable declare
	b1.Title = "An Introduction Of Programming In Go"
	b1.Author = "CALEB DOXY"
	b1.ISBN = "978-2236868969"
	b1.Price = 100.50
	b1.Page = 112

    fmt.Println(b1)// all value 
	fmt.Println(b1.Title)//separate value
	fmt.Println(b1.Author)
	fmt.Println(b1.ISBN)
	fmt.Println(b1.Price)
	fmt.Println(b1.Page)
}
______________________________________________________

package main 

import "fmt"

func main(){

b1 := struct{//annoymous struct
title string
author string
isbn string
price float32
page int
}{
title: "an introduction of progremming of book",
author: "CALEB DOXY",
isbn: "749-1715496505",
price: 100.50,
page: 112,
}
fmt.Println(b1)

}
