package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	switch r.Method {
		case "GET":		
			fmt.Fprintf(w, "Hello World! (Version Info: %s, Build Date: %s)", os.Getenv("VERSION_INFO"), os.Getenv("BUILD_DATE"))
			 http.ServeFile(w, r, "index.html")
		default:	
			fmt.Fprintf(w, "Unsupported.")
		}
}

func main() {
	http.HandleFunc("/", mainHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
