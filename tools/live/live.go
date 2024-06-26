package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		panic("No argument provided")
	}
	folderToServe := argsWithoutProg[0]
	fmt.Println("Serving folder: " + folderToServe)

	server:= http.NewServeMux()
	server.Handle("/", http.FileServer(http.Dir(folderToServe)))
	port := ":1313"
	fmt.Println("Server is running at http://localhost" + port + "/")
	http.ListenAndServe(port, server)
}




