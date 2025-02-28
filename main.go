package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	// fmt.Println("Hello World")
	fileServer := http.FileServer(http.Dir("./static"));

	http.Handle("/", fileServer);
	http.HandleFunc("/hello", helloHandler);
	http.HandleFunc("/form",formHandler);;


	fmt.Println("Server Started at Port 8080");

	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err);
	}

}


func helloHandler(w http.ResponseWriter ,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 Page not found ", http.StatusNotFound);
		return
	}

	if r.Method != "GET" {
		http.Error(w, "GET Method is not supported you should use GET Method", http.StatusMethodNotAllowed);
		return
	}

	fmt.Fprintf(w, "Hello Ebery Nyan \n");
	log.Println("Hello Requested \n");
}


func formHandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/form"{
		http.Error(w, "404 Page not found check your router ", http.StatusNotFound);
		return
	}

	if r.Method != "POST" {
		http.Error(w, "POST Method is not supported you should use POST Method", http.StatusMethodNotAllowed);
		return
	}

	if err := r.ParseForm(); err != nil{
		http.Error(w, "Error parsing form", http.StatusBadRequest);
		return
	}
	
	fmt.Fprintf(w, "POST request successful \n \n \n");
	name:= r.FormValue("name");
	address:= r.FormValue("address");
	fmt.Fprintf(w, "Name: %s \n", name);
	fmt.Fprintf(w, "Address: %s \n", address);


}