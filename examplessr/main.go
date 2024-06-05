package main

import (
	"fmt"
	"net/http"
	* "github.com/fengdotdev/coipossr/components"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := Text("Hello, World! This is a example coipossr app!")
		w.Write([]byte("Hello, World! This is a example coipossr app!"))
	})

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", server)

}
