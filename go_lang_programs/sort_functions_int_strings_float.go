package main
import (
	"fmt"
	"sort"
)

func main(){
	//Int sort functions
	nums:=[]int{50,90, 30, 10, 50}
	fmt.Println(nums)
	if sort.IntsAreSorted(nums) == false{
		sort.Ints(nums)
	}
	fmt.Println(nums)
	fmt.Println("Enter searched no:")
	var i int
	fmt.Scanln(&i)
	fmt.Println("at position :",sort.SearchInts(nums,i))

	//string sort functions
	str := []string {"US","UK","Germany","Australia","Japan"}
	fmt.Println(str)
	if sort.StringsAreSorted(str) == false{
		sort.Strings(str)
	}
	fmt.Println(str)
	fmt.Println("Enter searched string:")
	var s string
	fmt.Scanln(&s)
	fmt.Println("at position :",sort.SearchStrings(str,s))

	//float64 sort functions
	val:=[] float64 {2.5,6.3,1.5,9.8,4.7,1.1,5.2}
	fmt.Println(val)
	if sort.Float64sAreSorted(val)==false{
		sort.Float64s(val)
	}
	fmt.Println(val)
	fmt.Println("Enter searched val:")
	var f float64
	fmt.Scanln(&f)
	fmt.Println("at position :",sort.SearchFloat64s(val,f))





}