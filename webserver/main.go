package main

import (
	"fmt"
	"log"
	"net/http"	
)


func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(res, "hello")
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(res, "POST request succeful")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "Name -%s\n", name)
	fmt.Fprintf(res, "Address -%s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}