package main
import "fmt"
func main() {
	//1. declare variable using var keyword
	//variable auto knows datatypes 
	var a = 5
	var b = "pavan"
	var c = 34.23
	var d = true

	fmt.Printf(" a= %T "+" b= %T c = %T"+"d = %T\n",a,b,c,d)

	//variable default values
	var e uint16
	var f float64
	var g string
	var h bool

	fmt.Printf(" e = %d f = %f "+"g = %s h = %v \n",e,f,g,h)

	//declare multiple variables of same datatype
	var i,j,k uint8 = 2,3,4
	var l,m,n float32 = 2.3,4.5,5.6

	fmt.Println(i,j,k,l,m,n)

	//declare multiple variables of different datatype

	var o,p,q = 1,"pavan",2.3
	//fmt.Printf("%d %s %f",o,p,q)
	fmt.Println(o,p,q)

	//2. declare variable using short variable declaration
	//short variable declaration used to store local variables
	r:= 5
	s:="pavan"
	t:=45.34

	fmt.Println(r,s,t)

	//short variable declaration used to store multiple value
	u,v,w:=4,"kumar",34.123
	fmt.Println(u,v,w)

	// short variable declaration used to reinstallize already existing
	//variable  by declaring new variable in same lexiacal block
	x,y:=4,"pavan"
	fmt.Println(x,y)
	z,y:=6.3,"kumar"
	fmt.Println(z,y)




}
