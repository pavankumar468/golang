package main
import "fmt"

func main(){
	var rows int
	var k int = 1;
	fmt.Println("Enter rows:")
	fmt.Scanln(&rows)

	for i:=1;i<rows;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d ",k)
			k++
		}
		fmt.Println(" ")
	}
}