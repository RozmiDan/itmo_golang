package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRoot(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}

func handleTest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, r.URL.String())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/test/", handleTest)

	server := http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
	}

	fmt.Println("Started the server")
	
	server.ListenAndServe()

}