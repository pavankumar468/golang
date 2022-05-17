package main
import "fmt"

func main(){
	var rad float64
	var area float64
	const PI float64 = 3.14
	var circum float64

	PI64:=float64(PI)
	fmt.Printf("Enter radius:")
	//fmt.Scanf("%f",&rad)
	fmt.Scanln(&rad)
	area =PI64 *rad*rad
	fmt.Println("area of circle:",area)
	circum= 2* PI64 *rad
	fmt.Println("circumference of circle:",circum)
}