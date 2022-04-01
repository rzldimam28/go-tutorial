package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "CAN NOT FIND PAGE :(", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(writer, "HELLO ^^")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println("Could not load form. error:", err)
		return
	}
	fmt.Fprintln(writer, "Success Post Form Value")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintln(writer, "Name:", name)
	fmt.Fprintln(writer, "Address:", address)
}