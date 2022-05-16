package main

import (
	"fmt"

)
func main(){
	p:=34
	q:=20
	//Arithematic operator
	fmt.Println("Arithematic operator")
	result1:=p+q
	fmt.Printf("Result of p + q = %d ",result1)
	result2:=p-q
	fmt.Printf("\nResult of p + q = %d ",result2)
	result3:=p*q
	fmt.Printf("\nResult of p + q = %d ",result3)
	var result4 int = p/q
	fmt.Printf("\nResult of p + q = %d ",result4)
	var result5 =p%q
	fmt.Printf("\nResult of p + q = %d ",result5)

	fmt.Println("Relational operator")
	result6:=p==q
	fmt.Printf("Result of p + q = %v ",result6)
	result7:=p!=q
	fmt.Printf("\nResult of p + q = %v ",result7)
	result8:= p>q
	fmt.Printf("\nResult of p + q = %v ",result8)
	var result9 bool = p<q
	fmt.Printf("\nResult of p + q = %v ",result9)
	var result10 = p>=q
	fmt.Printf("\nResult of p + q = %v ",result10)
	var result11 = p<=q
	fmt.Printf("\nResult of p + q = %v ",result11)

	fmt.Println("\nlogical operator")
	var r int =23
	var s = 60
	if(r!=s && r<=s){
		fmt.Println("true")
	}
	 if(p!=q || p<=q){
		fmt.Println("true")
	}
	 if(!(p==q)){
		fmt.Println("false")
	}

	fmt.Println("\nBitwise operator")
	result12:=p&q
	fmt.Printf("Result of p & q = %d ",result12)
	result13:=p|q
	fmt.Printf("\nResult of p | q = %d ",result13)
	result14:=p^q
	fmt.Printf("\nResult of p ^ q = %d ",result14)
	var result15 = p>>1
	fmt.Printf("\nResult of p >> q = %d ",result15)
	var result16 =p<<1
	fmt.Printf("\nResult of p << q = %d ",result16)
	var result17 =p&^q
	fmt.Printf("\nResult of p &^ q = %d ",result17)

	fmt.Println("\n Assignment operator")
	var f int = 45
	var g int = 50

	f=g
	fmt.Println(f)
	f+=g
	fmt.Println(f)
	f-=g
	fmt.Println(f)
	f*=g
	fmt.Println(f)
	f/=g
	fmt.Println(f)
	f%=g
	fmt.Println(f)

	fmt.Println("\n Misc operator")
	//&: address of variable
	//*: pointer points to variable
	a:=4
	fmt.Println(a)
	fmt.Println("address of a",&a)

	b:=&a
	fmt.Println("value of b",b)
	fmt.Println("b points to",*b)



	


}