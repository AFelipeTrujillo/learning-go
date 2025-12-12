/*
* https://www.youtube.com/watch?v=jFfo23yIWac
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "POST request successful\n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprint(w, "Name = ", name, "\n", "Address = ", address)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Check that the URL path is /hello
		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		// Check that the request method is GET
		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		// Fprintf writes formatted output to w
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
