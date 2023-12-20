package main

import (
	"net/http"
	"io"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "WE ARE LIVE !!!")
	})
	
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "API IS LIVE !!!")
	})
	
	http.ListenAndServe(":8080", nil)
}