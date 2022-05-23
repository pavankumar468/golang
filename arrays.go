package main
import "fmt"

func main(){
	//declare variable using var keyword
	var arr[5] string 
	//{"ads","gfg","pavan","kum","java"}
	arr[0] = "ads"
	arr[1] = "gfg"
	arr[2] = "pavan"
	arr[3] = "kum"
	arr[4] = "java"

	//for i:=0;i<len(arr);i++{
		for _,v :=range arr{
		fmt.Println(v)
	}
	fmt.Println(len(arr))
	//declare variable using short variable declaration
	arr1 := [5]int {5,4,3,2,1}
	//declare variable using short variable declaration with variable no of arguments
	arr2 := [...]float64 {2.3,3.4,5.6,6.6,7.8}
	for _,v :=range arr1{
		fmt.Println(v)
	}
	fmt.Println(len(arr1))
	for _,v :=range arr2{
		fmt.Println(v)
	}

	//multi dimentional array
	arr3 :=[2][2] int {{12,23},{13,34}}

	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Printf("%d ",arr3[i][j])
		}
		fmt.Println()
	}

	//compare 2 arrays are equal or not
	arr4:=[5]int{5,4,3,2,1}
	fmt.Println(arr4==arr1)

	//default value of array is zero
	var arr5[5] int
	fmt.Println(arr5)
}