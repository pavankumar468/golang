package main

import (
	"fmt"
)
var a int = 100   // Global varibles are declared out block/functions/loops

func main(){
	var b int =34  // local varibles are declared out block/functions/loops
	fmt.Println(b)
	fmt.Println(a)
	a:=23
	fmt.Println(a)// if Global & local varibles having same name always perference given to local variable
	display()
}

func display(){
	//fmt.Println(b)
	fmt.Println(a)
}