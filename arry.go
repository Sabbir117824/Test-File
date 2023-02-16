LONG STRUCTRE

package main 

import "fmt"

func main(){
var students [3] string
students[0]= "sabbir"
students[1]= "sajjad"
students[2]= "shanto"

fmt.Println(students)
//fmt.Println(len(student))

}



SHORT STRUCTRE

package main 

import "fmt"

func main(){
 student := [...] string{"sabbir","sajjad","shanto"}

fmt.Println(student)
//fmt.Println(len(student))

}