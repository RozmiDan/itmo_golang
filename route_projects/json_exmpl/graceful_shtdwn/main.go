package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, req *http.Request){
		time.Sleep(time.Second * 2)
		w.Write([]byte("Hello"))
	})

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	go func(){
		err := server.ListenAndServe()
		if err != http.ErrServerClosed{
			log.Fatal(err)
		}
		fmt.Println("Server stopped")
	}()
	
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	termSig := <-termChan
	log.Println("Graceful shutdown started with signal", termSig)

	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := server.Shutdown(closeCtx)
	if err != nil{
		log.Println("Server shutdown failed with error:", err)
	}

	log.Println("Graceful shutdown completed")
}