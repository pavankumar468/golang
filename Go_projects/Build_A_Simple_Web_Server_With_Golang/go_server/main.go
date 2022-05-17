package main
import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:=r.ParseForm(); err!=nil{// we want people to submit something on form.html file ,it will POST request that will be parse form
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return 
	} 
	fmt.Fprintf(w,"POST request successful")
	// now filling form
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"Name = %s \n",name)
	fmt.Fprintf(w, "Address = %s\n",address)
}
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{  // if request path is not /hello then error
		http.Error(w,"404 not found",http.StatusNotFound)
		return 
	}
	if r.Method != "GET"{// when we type /hello in broswer by default is GET method & print hello, we dont want to POST anything to server. 
		http.Error(w,"method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello_pavan!") //print to the stream

}

func main(){
	fileServer := http.FileServer(http.Dir("./static")) //tells FileServer check to static directory& go_lang always 1st goto index.html
	http.Handle("/",fileServer) //start handling "/(root route i.e index.html)" by sending it fileserver
	http.HandleFunc("/form",formHandler)  //start handling "/form route " to form handler function
	http.HandleFunc("/hello",helloHandler)  //start handling "/hello route " to form handler function

	fmt.Printf("Starting server at port 8080\n ")
	if err:=http.ListenAndServe(":8080",nil); err !=nil{
		log.Fatal(err)
	}
}