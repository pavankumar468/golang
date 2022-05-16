package main

import "fmt"

func main() {
	//integer datatypes
	var x uint8 = 225
	fmt.Println(x, x-3)

	var y int16 = 23432
	fmt.Println(y+2, y-2)
	fmt.Printf("%T, %d\n",x,y)

	//float datatypes
	var f float32 = 45.56
	var g1 float64 = 3445.234
	fmt.Println(f, g1)
	
	a:=23.45
	b:=34.32
	c:=a-b
	fmt.Println("Result is %T: %f",c)
	fmt.Printf("%T\n",c)

	// string
	var str string = "pavan kumar"
	var str1 string = "hello"
	var str2 string = "pavan kumar"
	result1:= str==str1
	result2:= str==str2

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Printf("type of result1 id %T"+" type of result2 is %T",result1 , result2)
	fmt.printf("Length of string = %d " len(str))

	fmt.Println(str)

	// bool
	var b8 bool = true
	 fmt.Println(b8)
}
