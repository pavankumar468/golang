package main
import "fmt"

func rect(c,d int) (int,int){
	//var area,circum int
	area =2*(c+d)
	circum = c*d

	return area,circum
}

func main(){
	var a,b int 
	fmt.Println("Enter a,b :")
	fmt.Scanln(&a,&b)
	var area,circum int = rect(a,b)
	fmt.Printf("area of rectangle : %d\ncircum of rectangle: %d ",area,circum)
	

}