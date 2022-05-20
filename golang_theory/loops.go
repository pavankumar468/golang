//go lang  is having only one 
package main
import "fmt"

func main(){
	fmt.Println("1.simple use of for loop")
	//like "for" loop known no of iterations
	for i:=1;i<10;i++{
		fmt.Println(i)
	}
	//like "while " loop doesn't no of iterations
	i:=1
	for i<5{
		fmt.Println(i)
		i++
	}

	fmt.Println("2.Infinite loop")
	g:=5
	for{
		fmt.Println("pavan")
		if(g==10){
			break
		}
		g++
	}

	fmt.Println("3.Ranges in for loop")
	variable := [] string {"GFG", "Geeks", "GeeksforGeeks"}
	for i,j := range variable{
		fmt.Println(i,j)
	}

	for range "HELLo"{  // loop iterates depends on no of characters
		fmt.Println("hello")
	}
	fmt.Println("4.iterate over map using Ranges in for loop")

	pmap := map[int]string {
		22:"pavan",
		33:"kumar",
		44:"lol",
	}
	for i,j:=range pmap{
		fmt.Println(i,j)
	}

	fmt.Println("5.iterate over channel using Ranges in for loop")

	channel := make(chan int)
	go func(){
		channel <- 100
		channel <- 1002
		channel <- 1023
		channel <- 100000
		channel <- 134324
		close(channel)
	}()

	for i :=range channel{
		fmt.Println(i)
	}
	fmt.Println("5.iterate over Unicode code point for a string")

	for i,j := range "pavan kumar"{
		fmt.Println("the Index of %U is %d",j,i)
	}

	


}