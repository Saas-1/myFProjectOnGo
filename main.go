package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandle(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "request successful<br>")
	name := r.FormValue("name")
	address := r.FormValue("address")
	if name == "" || address == "" {
		http.Error(w, "Both fields are required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address=%s", address)

}

func helloHandle(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {

	fileserver := http.FileServer(http.Dir("./src/static"))
	http.Handle("/", fileserver)

	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", helloHandle)
	//	http.HandleFunc("/tictactoe", tictactoeHandle)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
