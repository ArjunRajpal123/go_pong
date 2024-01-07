package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/arjunrajpal123/go_pong/pages/landing_page"
)

func main() {
	component := landing_page.Content("World") // Use the correct package name when calling the function
	http.Handle("/", templ.Handler(component))

	// Serve static files in the "js" directory
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	// Serve static files in the "css" directory
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
