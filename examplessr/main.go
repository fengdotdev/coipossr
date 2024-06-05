package main

import (
	"fmt"
	"net/http"

	. "github.com/fengdotdev/coipossr/components"
	"github.com/fengdotdev/coipossr/render"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := render.NewPage(w, r, Text("Hello, SSR!"))
		page.Render()
	})

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", server)

}
