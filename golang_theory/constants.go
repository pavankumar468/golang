package main
import "fmt"
const PI =3.14
const A = "GFG"
var B = "Geeksforgeeks"
const v = false
func main(){
	fmt.Printf("hello %v",PI)
	var helloworld = A+"  "+B
	helloworld+="!"
	fmt.Println(helloworld)
	fmt.Println(A=="GFG")
	fmt.Println(B<A)

	value:=v
	fmt.Println(value)
}