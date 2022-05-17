package main
import "fmt"

func palidrome(no uint64) (bool){
	var r uint64
	var sum uint64
	var temp uint64 = no
	for no!=0 {
		r=no%10
		sum =sum*10+r
		no=no/10
	}
	if temp == sum{
		return true
	}else{
		return false
	}
}

func main(){
	var no uint64
	fmt.Println("Enter a no:")
	fmt.Scanln(&no)
	fmt.Println("Is Palidrome = ",palidrome(no))

}