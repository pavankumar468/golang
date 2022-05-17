package main
import "fmt"

func main(){
	var x int
	fmt.Println("Enter a no:")
	fmt.Scanln(&x)

	for i:=1;i<=10;i++{
		fmt.Printf("%d * %d = %d\n",x,i,x*i)
	}
}