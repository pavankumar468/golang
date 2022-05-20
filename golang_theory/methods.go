//methods are lly to function but main difference is methods takes receiver(either struct or nonstruct) as input and can use receiver properties
package main
import "fmt"

/*
fmt.Println("1.Method with struct type receiver")

type emp struct{
	name string
	age int8
	sal float64
}

func (e emp)show(){
	fmt.Println(e.name, e.age, e.sal)
}

func main(){

	result:= emp{name:"pavan",age:23,sal:15.23}
	result.show()
}*/

//fmt.Println("2.Method with Non-Struct Type Receive")
/*type data int

func (d1 data)Multiple(d2 data) data{
	return d1*d2
}

func main(){
	value1:=data(23)
	value2:=data(34)
	res:=value1.Multiple(value2)
	fmt.Println(res)
}*/

//methods with pointer receiver , with the help of pointer receiver the changes made in method will affect on caller, which will be not affected 
// value receiver methods

type emp struct{
	name string 
	age int
	sal float64
}

func (e *emp)show(){
 fmt.Println((*e).name, (*e).age, (*e).sal)
}

func (e *emp) change(chsal float64){
		(*e).sal = chsal
}
func (f emp) valuechange(ename string){
	f.name = ename
	fmt.Println(f.name)
}
func main(){
	res := emp{"pavan",8,12.23}
	res.show()
	fmt.Println(res.name,res.age,res.sal)

	res.change(45.6723)
	res.show()

	p:=&res
	p.change(23.34)
	res.show()

	q:=&res
	q.valuechange("sai kumar")
	fmt.Println(q.name) 
	res.show()
}
