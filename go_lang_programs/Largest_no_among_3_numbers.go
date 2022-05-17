package main
import "fmt"

func main(){
	var a,b,c int
	fmt.Println("Enter a,b,c :")
	fmt.Scanf("%d %d %d",&a,&b,&c)

	if a>b && a>c{
		fmt.Println("a is big")
	}else if b>a && b>c{
		fmt.Println("b is big")
	}else{
		fmt.Println("c is big")
	}


	fmt.Println(a,b,c)
}