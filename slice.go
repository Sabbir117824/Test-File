package main

import "fmt"
//import "reflect"//for know the data type

func main(){
	var fruits []string

	fruits = append(fruits,"Apple","banna","mango")

	fmt.Println(fruits)
	fmt.Printf("%T",fruits)// for know the varrible type
	fmt.Println(fruits, len(fruits))// for know the lenth of value
	fmt.Printf("%T\n",fruits)//for creat a new line 
	//a := reflect.TypeOf(fruits).Kind().String()//another way to know data type

	 //fmt.Println(a)

}