package main

import (
	"log"
	"net/http"
	"time"
)

func adminMiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("middleware admin started")
		w.Write([]byte("Hello, admin"))
		next.ServeHTTP(w,r)
		log.Println("middleware admin finished")
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("accessLogMiddleware", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("panicMiddleware started")
		defer func(){
			if err := recover(); err != nil {
				log.Println("Panic canceled")
				http.Error(w, "Server error", 500)
			}
		}()
		next.ServeHTTP(w,r)
		log.Println("panicMiddleware finished")
	})
}

func panicPage(w http.ResponseWriter, r *http.Request) {
	panic("this must me recovered")
}

func startedPageFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Started handler")
	w.Write([]byte("Hello"))
	time.Sleep(time.Millisecond * 10)
	log.Println("Ended handler")
}

func loginFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("login handler started")
	w.Write([]byte("You on login page"))
	log.Println("login handler finished")
}

func logoutFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("logout handler started")
	w.Write([]byte("You on login page"))
	log.Println("logout handler finished")
}

func adminFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("logout handler started")
	w.Write([]byte("You on login page"))
	log.Println("logout handler finished")
}

func main() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminFunc)
	adminMux.HandleFunc("/admin/panic", panicPage)
	adminHandler := adminMiddlewareFunc(adminMux)

	accessMux := http.NewServeMux()
	accessMux.HandleFunc("/login", loginFunc)
	accessMux.HandleFunc("/logout", logoutFunc)
	accessMux.HandleFunc("/", startedPageFunc)
	accessMux.Handle("/admin/", adminHandler)
	
	accessHandler := accessLogMiddleware(accessMux)
	accessHandler = panicMiddleware(accessHandler)

	server := http.Server{
		Addr: ":8080",
		Handler: accessHandler,
	}

	server.ListenAndServe()
}