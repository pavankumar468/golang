package main
import "fmt"

func myfunec(a [] int, n int)(int){
	var sum int = 0
	for i:=0;i<n;i++{
		sum+=a[i]
		fmt.Printf("%d ",a[i])
	}
	return sum
}

func main(){
	fmt.Println("copy array by value, if changes done in original array , doesnt affect copied array")
	arr:=[5]int {2,3,4,5,6}
	arr1 := arr
	fmt.Println("original array")
	for _,v:=range arr{
		fmt.Printf("%d ",v)
	}

	arr[2] = 56
	fmt.Println("\noriginal after changes array")
	for _,v:=range arr{
		fmt.Printf("%d ",v)
	}
	fmt.Println("\ncopied array")
	for _,v:=range arr1{
		fmt.Printf("%d ",v)
	}
    fmt.Println("\n---------------------------------------")
	fmt.Println("copy array by reference, if changes done in original array , affect copied array")

	arr2 := &arr

	fmt.Println("original array")
	for _,v:=range arr{
		fmt.Printf("%d ",v)
	}

	arr[1] = 898
	fmt.Println("\ncopied array")
	for _,v:=range arr2{
		fmt.Printf("%d ",v)
	}
	arr3:=[...]int{3,5,6,7,8,23,5}

	fmt.Println("Passing array into function")
	 sum1:= myfunec(arr3,len(arr3))
	 fmt.Println(sum1)




}