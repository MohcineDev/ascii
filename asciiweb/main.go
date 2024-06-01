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
		http.Error(writer, "Error parsing the file ", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(writer, name, nil)
	if err != nil {
		fmt.Println("Error when executing the template", err)
	}
}
func handleFunc(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fileName := "index.html"
		parseTheFile(fileName, res)

	case "/about":
		fmt.Fprint(res, "about")
	default:
		fileName := "404.html"
		parseTheFile(fileName, res)
	}
	fmt.Printf(req.Method)
}

func main() {
	const PORT = "8080"

	http.HandleFunc("/", handleFunc)

	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
