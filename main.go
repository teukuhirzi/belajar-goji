package main

import (
	"fmt"
	"net/http"
	//"goji.io"
	//"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//name := pat.Param(r, "name")
	//fmt.Fprintf(w, "Hello, %s!", name)
	fmt.Println("Hello World")
}

func main() {
	//mux := goji.NewMux()
	//mux.HandleFunc(pat.Get("/hello/:name"), hello)

	//http.ListenAndServe("localhost:8000", mux)
}
