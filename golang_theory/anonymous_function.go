package main
import (
	"fmt"
	"math"
	//"strings"
)

func myfunc(a  func(int,int) (int,int))  (area int, circum int){
	area, circum = a(6,9)
	return 
}

func myfunc1( a func(...int) int) (result int){
	str :=[]int {23,345,23,14}
	result = a(str...)
	return
}


func main(){
	fmt.Println("1.no parameter no returntype anonymous fun")
	func(){
		fmt.Println("Hello pavan")
	}()

	fmt.Println("2.with parameter anonymous fun")

	func(i int) {
		fmt.Println("Hello pavan",i)
	}(10)

	fmt.Println("3.func address given to variable")

	Add:= func (i,j int) (area,circum int){
			area = i*j
			circum = int(2* math.Pow(float64(i),float64(2)))
			return 
			}
		area,circum :=Add(4,5)
		fmt.Println(area,circum)
		fmt.Printf("Type of %T",Add)

	fmt.Println("4.call backs")
	area1, circum1 :=myfunc(Add)
	fmt.Println(area1,circum1)

	fmt.Println("4.call backs for strings")

	take:= func (ele... int) (total int){
		//total=""
		for _,v:= range ele{
			//total = total +" "+ v
			total+=v
		}
		return 
	}

	
	//str1 = [] string{"anonymous","function","from","another"}
	str := []int {1,4,6,3}
	 
	var word int = take(str...)
	fmt.Println(word)
	fmt.Printf("type = %T\n",take)
	result := myfunc1(take)
	fmt.Println(result)

}