package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func parseTheFile(name string, writer http.ResponseWriter) {
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		fmt.Println("Error parsing the file", err)
		return
	}
	err = tmpl.ExecuteTemplate(writer, name, nil)
	if err != nil {
		fmt.Println("Error when exxecuting the template", err)
	}
}
func hundleFunc(writer http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fileName := "index.html"
		parseTheFile(fileName, writer)
		/*		temp1 := template.Must(template.ParseFiles("index.html"))

				films := map[string][]Film{

					"Films": {
						{Title: "The GodFather", Director: "Francis Ford Coppala"},
						{Title: "Blade Runner", Director: "Ridley scott"},
						{Title: "The thing", Director: "John Carpenter"},
					},
				}
				temp1.Execute(w, films)*/
	case "/about":
		fmt.Fprint(writer, "about")
	default:
		fileName := "404.html"
		parseTheFile(fileName, writer)
	}
	fmt.Printf(r.Method)
}

func main() {
	const PORT = "8080"
	fmt.Printf("running on %v...", PORT)

	http.HandleFunc("/", hundleFunc)

	http.ListenAndServe(":"+PORT, nil)
}
