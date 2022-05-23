package main
import "fmt"
//slice components pointers,length,capacity
//pointer: position pointing to
//length: length of slice
//capacity: maximum length it can expand

func main(){
	fmt.Println("declare slice")
	//create slice using var keyword
	var sli = []string{"Geeks", "for", "Geeks", "GFG"}
	fmt.Println(sli)
	//create slice using short variable declaration
	sli1 := []int{2,4,6,5,8} 
	fmt.Println(sli1) 
	//sli1[5]=90  //we cannot add element into slice  with "="
	sli1=append(sli1,34,23,12,5)

	fmt.Println(sli1)

	fmt.Println("create slice from array")
	ca:=sli[1:] //print from position 1 to end
	a:=sli1[2:5] //print from position 2to 4
	b:=sli1[:4]
	d:=sli1[:]
	fmt.Println(a,b,ca,d)
	sli2:=[4]int{}
	sli2[0] = 2
	sli2[1] = 4
	sli2[2] = 56
	sli2[3] = 234
	//fmt.Printf("Type =  %T",sli2)
	fmt.Println(sli2)

	fmt.Println("dynamically create slice")
	score := make([]string,4,6) //make(datatype,length,capacity)

	score[0]="pavan"
	score[1]="asd"
	score[2]="fgf"
	score[3]="ghd"
	//score[4]="nbv" out of range can be added using append keyword
	fmt.Println(score)
	score = append(score, "afed","bggfs","bn b","ghh","fgsdf")
	fmt.Println(score)





}