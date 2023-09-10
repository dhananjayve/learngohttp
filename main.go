package main

import (
	"fmt"
	"net/http"
	"log"
	//"time"
)



// don't touch below this line

type request struct {
	path string
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World test"))
}

func middlewareCors(next http.Handler) http.Handler {
	fmt.Println("serving request")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			 return
		}
		
		// w.Write([]byte("Hello World test"))
		next.ServeHTTP(w, r)
		
	})
}

func main() {
	fmt.Println("Starting")
	mux := http.NewServeMux()
	// mux.HandleFunc("/", HelloHandler)
	// http.ListenAndServe("localhost:8080", mux)
	corMux:= middlewareCors(mux)
	srv := &http.Server{
		Addr: ":8080",
		Handler: corMux,
	}
	log.Println("serving on post : 8080")
	log.Fatal(srv.ListenAndServe())
}

