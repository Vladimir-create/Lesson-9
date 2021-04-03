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
		io.WriteString(w, "successful post")	
	/*if req.Method == "GET" {
		io.WriteString(w, `<a href="/page1">GOTO PAGE1</a>`)
		//w.Write()
	} 
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")*/
	} else {
		w.WriteHeader(405)
	}	
}

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
