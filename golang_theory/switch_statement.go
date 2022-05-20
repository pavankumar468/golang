package main

import "fmt"

// switch statement used to transfer execution to different blocks of code based on value

func main(){
			//Expression switch
	fmt.Println("switch statement with option statement")
	switch day:=2; day{
	case 1: fmt.Println("Monday") 
	case 2: fmt.Println("Tuesday")
	case 3: fmt.Println("Wednesday") 
	case 4: fmt.Println("Thursday") 

	default:  fmt.Println("Invalid")
	}
/*
	fmt.Println("switch statement without option statement")
	var value int =1

	switch{
	case value == 1: fmt.Println("Hello")
	case value ==2:  fmt.Println("Bonjour")
	default : fmt.Println("Invalid")
	}*/
		
	fmt.Println("type switch")	//type switch
	var value interface{}
	switch q:=value.(type){
	case bool:  fmt.Println("value is of boolean type")
	case uint16 : fmt.Println("value is of uint16 type")
	case int: fmt.Println("value is of int type")
	default: fmt.Println("value is of type %T",q)
	}



}