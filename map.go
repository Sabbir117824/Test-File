package main

import "fmt"

func main(){
	x := make(map[string]string)
	x["name"] = "Sabbir"
	x["hight"] = "5.6"
	x["address"] = "Chittagong"
	//delete (x,"adress")// for delete the varrible

	fmt.Println(x)

}