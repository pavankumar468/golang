package main
import "fmt"

type salary func(int,int) int
type Emp struct{
	name,city string
	age int
	//Function as field
	sal salary

	//function as anonymous function
	pend func(int,int) int
}
func main(){
	result :=Emp{
		name: "pavan",
		city: "Ydf",
		age: 23,
		sal: func(ma int, ja int) int{
			return ma*ja+2
		},
	}

	result1:=Emp{
		name:"kumar",
		city:"hyd",
		age: 34,
		pend:func(a int,b int) int{
			return a*b+8
		},
	}
	fmt.Println(result.name,result.age,result.sal(2,8))
	fmt.Println(result1.pend(2,7))
}