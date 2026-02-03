package main

import "net/http"

func helloHandler(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("Hello, World!"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("About Page"))
}

func main() {
 mux := http.NewServeMux()
 mux.HandleFunc("/hello", helloHandler)
 mux.HandleFunc("/about", aboutHandler)

 err := http.ListenAndServe(":3000", mux)
 if err != nil {
  panic(err)
 }
}