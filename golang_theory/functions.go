package main
import (
	"fmt"
	"strings"
)

func area(c,d int) int{
	result:=c*d+1
	return result
}

func swap_by_value(c,d int) {
	temp:= c
	c = d
	d=temp
}

func swap_by_reference(c,d *int){
	temp:= *c
	*c = *d
	*d = temp
}

func join_str(elements...string) string{
	return strings.Join(elements,"@")
}

func mult(nums...int) int{
	total := 1
	for _,a:= range nums {
		total *= a
	}
	return total
}

func myfunc(p,q,r uint16) (int,int,int){
	return int(p+q),int(q+r),int(p*q*r)
}

func calc(j,k int) (rect int,sq int){
	rect = j*k
	sq = j*j
	return 
}

func main(){
	var a int =4
	var b int =9

	fmt.Println("1.simple function")
	fmt.Println("area = ",area(a,b))

	fmt.Println("2.call by value")
	fmt.Println("before a,b:",a,b)
	swap_by_value(a,b)
	fmt.Println("after a,b:",a,b)

	fmt.Println("3.call by reference")
	fmt.Println("before a,b:",a,b)
	swap_by_reference(&a,&b)
	fmt.Println("after a,b:",a,b)

	fmt.Println("4.passing arguments to variadic func")
	//0 arguments
	fmt.Println(join_str())
	fmt.Println(join_str("GEEK", "GFG"))
	fmt.Println(join_str("Geeks", "for", "Geeks"))
	fmt.Println(join_str("G", "E", "E", "k", "S"))

	fmt.Println("5.passing slice to variadic func")
	value :=[]string {"geeks", "FOR", "geeks"}
	fmt.Println(join_str(value...))

	fmt.Println("6.passing int slice to variadic func")
	nums := []int {1,2,3,4,5}
	fmt.Println(mult(nums...))

	fmt.Println("7.Function Returning Multiple Values")
	var p,q,r uint16
	fmt.Println("Enter p,q,r:")
	fmt.Scanf("%v %v %v",&p,&q,&r)
	var v1,v2,v3 = myfunc(p,q,r)
	fmt.Printf("v1 =%d ,v2=%d, v3=%d",v1,v2,v3)

	fmt.Println("\n8.Function Returning Multiple Values with named return parameters")
	
	var rect,sq int = calc(7,9)
	fmt.Printf("rect = %d ,sq =%d",rect,sq)



}