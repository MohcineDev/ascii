package main

import (
	"asciiweb/ascii"
	"fmt"
	"html/template"
	"net/http"
)

type asciiS struct {
	data string
}

func parseAndExecute(name string, res http.ResponseWriter, btata string) {
	tmpl, err := template.ParseFiles(name)
	// Title := "Ascii Art Web Project"
	data := ascii.Generate(btata)
	fmt.Println("btata : ", btata)
	if err != nil {
		http.Error(res, "Error parsing the file ", http.StatusInternalServerError)
		return
	}
	if name == "ascii-art.html" {

		err = tmpl.ExecuteTemplate(res, name, data)
	} else {
		err = tmpl.Execute(res, nil)
	}

	if err != nil {
		fmt.Println("Error when executing the template", err)
	}
}
func handleFunc(res http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/":
		fileName := "index.html"

		parseAndExecute(fileName, res, "")

	case "/ascii-art":
		fileName := "ascii-art.html"
		var a asciiS

		a.data = req.FormValue("input")
		fmt.Println("qsd.data: ", a.data)

		parseAndExecute(fileName, res, a.data)
	default:
		fileName := "404.html"
		parseAndExecute(fileName, res, "")
	}
}

func main() {
	const PORT = "8000"

	fs := http.FileServer(http.Dir("styles/"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))

	http.HandleFunc("/", handleFunc)

	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
