package main
import ( 
	"fmt" 
	"reflect"
)

type Address struct{  //defining a struct type
	name,city string
	pin uint16
}

type Child struct{
	age int
	detail Address
}
/*type child2 struct{
	branch string
	ch child1
}*/
func main(){
	//Declaring a struct init all values with 0
	var a Address
	fmt.Println(a)

	//Declaring and init struct with string literal
	a1:=Address{"pavan","HYD",500}
	fmt.Println(a1)

	//Declaring & init structs with naming fields with string literals
	a2:=Address{name:"kumar",city:"skl",pin:120}
	fmt.Println(a2)

	//Uninitialized field set to 0
	a3:=Address{name:"sai"}
	fmt.Println(a3)

	//Accessing fields in structs
	fmt.Println(a1.name)
	fmt.Println(a1.pin)
	a1.pin = 890
	fmt.Println(a1)

	//Pointers to struct r used to assign memory address of one struct variable to another variable
	a4:=&Address{"tamm","mum",900}
	fmt.Println(a4)
	fmt.Println((*a4))
	fmt.Println((*a4).city)
	fmt.Println(a4.name)
	fmt.Println((a4).pin)

	a5 := &a1
	fmt.Println(a5)
	fmt.Println((*a5))
	fmt.Println((*a5).city)
	fmt.Println(a5.city)

	a6:=Address{name:"satya",city:"HYD1",pin:5000}
	a7:=Address{name:"satya",city:"HYD1",pin:5000}
	
	a8:=a6
	//Struct equality
	if a6 == a8{
		fmt.Println("Variable a1 is equal to variable a6")
	}else {
		fmt.Println("Variable a1 is not equal to variable a6")
	}

	fmt.Println("a1 == a6",reflect.DeepEqual(a6,a7))
	fmt.Println("a1 != a3",reflect.DeepEqual(a1,a3))\


}