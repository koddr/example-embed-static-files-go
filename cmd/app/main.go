package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/koddr/example-embed-static-files-go/internal/box"
)

// PageData structure
type PageData struct {
	Title       string
	Heading     string
	Description string
}

func main() {
	// Define embed file (for a short)
	index := string(box.Get("/index.html"))

	// Template
	tmpl := template.Must(template.New("").Parse(index))

	// Handle function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:       "The easiest way to embed static files into a binary file",
			Heading:     "This is easiest way",
			Description: "My life credo is 'If you can not use N, do not use N'.",
		}

		// Execute template with data
		if err := tmpl.Execute(w, data); err != nil {
			log.Fatal(err)
		}
	})

	// Start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
