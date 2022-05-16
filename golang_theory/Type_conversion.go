package main
import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	fmt.Println("Convert string to integer type ")
	//using Atoi function
	// func Atoi(s string) (int,error)
	strvar := "1999"
	intvar,err:= strconv.Atoi(strvar)
	fmt.Println(intvar,err,reflect.TypeOf(intvar))

	//using fmt.Sscan= used to scan string with spaces and store it in a variable
	/*strvar1 := "1999"
	intvalue:=0
	_ ,err := fmt.Sscan(strvar1, &intvalue)
	//_, err := fmt.Sscan(strvar1, &intvalue)
	fmt.Println(intvalue,err,reflect.TypeOf(intvalue))*/

	//using ParseInt = accept string with base(0,2,36) and bitsize(0to64) 
	//func ParseInt(s string, base int, bitSize int) (i int64, err, error)
	intvar1,err:=strconv.ParseInt(strvar,0,32)
	fmt.Println(intvar,err,reflect.TypeOf(intvar1))

	fmt.Println("Convert string to integer type ")
	// func ParseFloat(s String, bitSize int) (float64,error)
	s:="23.12341"
	intvar2,err := strconv.ParseFloat(s,32)
	fmt.Println(intvar,err,reflect.TypeOf(intvar2))

	fmt.Println("Convert string to bool type ")
	s1:="true"
	intvar3,err := strconv.ParseBool(s1)
	fmt.Println(intvar3,err,reflect.TypeOf(intvar3))

	fmt.Println("convert Boolean Type to String ")
	s2:=false
	intvar4:= strconv.FormatBool(s2)
	fmt.Println(intvar4,reflect.TypeOf(intvar4))

	fmt.Println("Convert Float to String type ")
	//func FormatFloat(f Float64, fmt byte, prec, bitSize int) string
	s3:=23.45654
	fmt.Println(s3,reflect.TypeOf(s3))
	str:=strconv.FormatFloat(s3,'E',-1,32)
	fmt.Println(str,reflect.TypeOf(str))

	fmt.Println("Convert Integer to String  type ")
	// using FormatInt
	// func FormatInt( i int64, base int) string
	var i int64 = 124
	fmt.Println(i,reflect.TypeOf(i))
	var x string = strconv.FormatInt(i,10)
	fmt.Println(x,reflect.TypeOf(x))

	//using fmt.Sprintf() used to convert any type to string
	//fmt.Sprintf() string

	x1:=fmt.Sprintf("%v",i)
	fmt.Println(x1,reflect.TypeOf(x1))

	fmt.Println("Convert Int data type to Int16 Int32 Int64 ")
	var t int =10
	fmt.Println(reflect.TypeOf(t))

	i16:=int16(i)
	fmt.Println(reflect.TypeOf(i16))

	i32:=int32(i)
	fmt.Println(reflect.TypeOf(i32))

	i64:=int64(i)
	fmt.Println(reflect.TypeOf(i64))

	fmt.Println("Convert Float32 to Float64 and Float64 to Float32")
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	f64:=float64(f32)
	fmt.Println(reflect.TypeOf(f64))

	var d64 float64=1097.655698798798
	d32:=float32(d64)
	fmt.Println(reflect.TypeOf(d32))

	fmt.Println("Converting Int data type to Float ")

	j32:=int32(f32)
	fmt.Println(reflect.TypeOf(j32))
	j64:=int64(f64)
	fmt.Println(reflect.TypeOf(j64))

	
}