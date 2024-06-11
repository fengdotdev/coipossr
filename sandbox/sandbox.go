package main

import (
	"fmt"
	"net/http"

	c "github.com/fengdotdev/coipossr"
	"github.com/fengdotdev/coipossr/builder"
	"github.com/fengdotdev/coipossr/components/txt"
)

func main() {
	Static()
}



func Static (){
	home := c.StaticPage(txt.Text("Hello, World!", txt.TextOptions{Color: "red"}))
	contact := c.StaticPage(txt.Text("Contact Us", txt.TextOptions{Color: "green"}))
	about := c.StaticPage(txt.Text("About Us", txt.TextOptions{Color: "blue"}))
	static := builder.NewStaticBuilder()
	static.AddPage("/", *home)  
	static.AddPage("/contact", *contact)
	static.AddPage("/about", *about)
	err := static.Build()
	if err != nil {
		fmt.Println(err)
	}
}


func SSR(){
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := c.Page(w, r, txt.Text("Hello, World!", txt.TextOptions{Color: "red"}))
		page.Render()
	})
	port := ":1313"
	fmt.Println("Server is running at http://localhost" + port + "/")
	http.ListenAndServe(port, server)
}