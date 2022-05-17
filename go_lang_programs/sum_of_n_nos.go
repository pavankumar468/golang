package main
import "fmt"

var no uint64

func sumof(no uint64) uint64{
	var i uint64
	var sum uint64

	for i=1; i<=no ;i++{
		sum+=i
	}
	 return sum
}

func main(){
	fmt.Println("Enter a no:")
	fmt.Scanln(&no)
	fmt.Println("sum = ",sumof(no))

}