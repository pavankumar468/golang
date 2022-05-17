package main
import "fmt"
var no uint16
var i int =1
var fact uint64=1

func factorial(n uint16) (uint64){

	if(n<0){
		fmt.Println("factorial of -ve nos doest exist")
	}else{
		for i<=5{
			fact *=uint64(i)
			i++
		}
	}
	return fact
}


func main(){
	fmt.Println("Enter a no")
	fmt.Scanln(&no)
	fmt.Println("Factorial is:",factorial(no))
}