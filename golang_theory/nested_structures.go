package main
import "fmt"

type Emp struct{
	name,city string
	age int
}

type Extra struct{
	sal float32
	branch string
	detail Emp
}

//Anonymous fields
type student struct{
	int
	float64
	string
}

//promoted fields
type Child struct{
	sub string
	Emp
}

func(d Emp) promotmethod(tage int) int{  //Emp struct method
	return d.age*tage
}

func (d Extra) promotmethod(ename string) string{
	 d.detail.name=ename
	 return  d.detail.name
} 


func main(){

	//Declaring nested structures
	result:=Extra{
		sal: 34.56,
		branch: "go",
		detail: Emp{name: "pavan",city:"HYD",age:123},
	}
	fmt.Println(result)

	//Anonymous struct
	result1:=struct{
		name string
		city string
		age int
	}{
		name: "pavan",city:"HYD",age:123,
	}

	fmt.Println("Anonymous struct used only once")
	fmt.Println(result1.name,result1.age,result1.city)

	value:=student{123,2345.13,"Bud"}
	fmt.Println("Anonymous fields")
	fmt.Println(value.int,value.float64,value.string)


	fmt.Println("Promoted Fields , types acts as fields")
	result2:=Child{
		sub: "English",
		Emp:Emp{name: "pavan",city:" krl",age: 45},
	}

	fmt.Println(result2.sub,result2.age,result2.name,result2.city)

	fmt.Println("Promoted methods , types acts as fields")

	result4:=Extra{
		sal: 23.234,
		branch: "ece",
		detail: Emp{"sam","hyd2",32432},
	}

	fmt.Println(result4.sal,result4.detail.name)
	fmt.Println("total age = ",result4.detail.promotmethod(23))
	fmt.Println("total age = ",result4.promotmethod("deepake"))


}