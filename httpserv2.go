package main

import (
	"net/http"
	"io"
	"fmt"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if (req.Method == "OPTIONS") {
		w.Header().Add("Access-Control-Allow-Origin", "*") 			
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET") 			
		w.Header().Add("Access-Control-Allow-Headers", "content-type, accept, accept-language, content-language") 
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }	
		fmt.Printf("Body is %s\n", data)
		io.WriteString(w, "successful option")
	}	
	if req.Method == "GET" {
		w.Header().Add("Access-Control-Allow-Origin", "*") 			
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET") 			
		w.Header().Add("Access-Control-Allow-Headers", "content-type, accept, accept-language, content-language") 
		io.WriteString(w, "successful get")
	} 
	if req.Method == "POST" {
		w.Header().Add("Access-Control-Allow-Origin", "*") 			
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET") 			
		w.Header().Add("Access-Control-Allow-Headers", "content-type, accept, accept-language, content-language") 
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else {
		w.WriteHeader(405)
	}	
}

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
