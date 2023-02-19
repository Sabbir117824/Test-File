/*package main// if else a fixed way
import "fmt"
func main(){
	fmt.Println("Enter your age:")
	var age int
	fmt.Scanf("%d", &age)
	if age < 3{
		fmt.Println("infent")
		
	}else if age>2 && age<13{
		fmt.Println("children")
	}else if age>12 && age<=19{
		fmt.Println("Teen age")
	}else {
		fmt.Println("adult")
	}
	
	fmt.Println(age)
}

package main//switch case 
import "fmt"
func main(){
	fmt.Println("Enter your age:")
	var age int
	fmt.Scanf("%d", &age)
	switch age {
	case 2:
		fmt.Println("infent")
		fallthrough//same case a ar por ar case follow korar jonno
	case 3,4,5,6,7,8,9,10:
		fmt.Println("childrean")
	case 11,12,13,14,15,16,17,18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
	
	
	fmt.Println(age)
}*/

/*package main 
import "fmt"
func main(){
	for i:=0; i<9; i++{// syntex for loop
		fmt.Println(i)
	}
}
package main 
import "fmt"
func main(){
	student := [] string {"SABBIR HOSSAN","SAJJAD","SHANTO"}// arry syntex 
	for i, stu := range student {// range loop syntex
		fmt.Println(i,stu)
	}
}
package main 
import "fmt"
func main(){
	i := 0
	for i<10{// while loop
		fmt.Println("MASTER ACADEMY")
		i++
	}
}*/