package main

import (
    "fmt"
    "net/http"
    "log"

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
	    fmt.Println("Hello, World!")
    },
    )
    http.ListenAndServe(":8000", nil)
    log.Print("The service is ready to listen and serve.")
    log.Fatal(http.ListenAndServe(":8000", nil))

}
