package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func parseTheFile(name string, res http.ResponseWriter) {
	tmpl, err := template.ParseFiles(name)
	Title := "Ascii Art Web Project"
	if err != nil {
		http.Error(res, "Error parsing the file ", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(res, name, Title)
	// err = tmpl.ExecuteTemplate(res, name, doc{"Ascii Art Web Project"})
	if err != nil {
		fmt.Println("Error when executing the template", err)
	}
}
func handleFunc(res http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/":
		fileName := "index.html"
		parseTheFile(fileName, res)

	case "/ascii-art":
		input := req.FormValue("input")
		fmt.Println("input : ", input)

		fmt.Fprint(res, "ascii-art")
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
