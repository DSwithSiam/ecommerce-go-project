package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Hello World")
}



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", HomeHandler)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server runtime error: " + err.Error())
	}


}