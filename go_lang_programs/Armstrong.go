package main
import (
	"fmt"
	"math"
)

func count(n uint64) (uint64){
	var i uint64=0
	for n!=0{
		i++
		n = n/10
	}
	return i
}

func main(){
	var no uint64
	var r float64
	var sum float64
	
	fmt.Println("Enter a no:")
	fmt.Scanln(&no)
	var temp,temp2 uint64 = no,no
	var c uint64 =count(no)

	for temp!=0{
		r = float64(temp%10)
		sum = sum+ math.Pow(r,float64(c))
		temp=temp/10
	}

	if(temp2 == uint64(sum)){
		fmt.Println("Armstrong no")
	}else{
		fmt.Println("not Armstrong no")
	}
	
}