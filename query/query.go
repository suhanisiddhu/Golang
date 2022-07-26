package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":8000", nil)
}
func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

/*
http.Handlefunc("/",server)
http.ListenAndServer(addr:":8000",handler:nil)
func serve(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w,"hello %s",r,URL.Path[1:])
*/
