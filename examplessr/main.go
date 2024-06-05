package main

import (
	"fmt"
	"net/http"

	. "github.com/fengdotdev/coipossr"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Page(w, r, Text("Hello World!"))
		page.Render()
	})

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", server)

}
