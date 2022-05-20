package main
import "fmt"
//Break is used to stop execution of loop
// continue is to stop current iteration and continue remaining iterations

func main(){
	for i:=1;i<50;i++{	
		//i*=2
		if(i==6||i==14||i==30){
			continue;
		}
		if(i==42||i==32){
			break;
		}
		
		fmt.Println(i)
	}

}