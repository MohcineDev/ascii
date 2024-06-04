package main

import (
	"asciiweb/ascii"
	"fmt"
	"html/template"
	"net/http"
)

type asciiS struct {
	input  string
	banner string
}

func parseAndExecute(name string, res http.ResponseWriter, btata asciiS) {
	tmpl, err := template.ParseFiles(name)
	// Title := "Ascii Art Web Project"
	data := ascii.Generate(btata.input, btata.banner)
	fmt.Println("btata : ", btata)
	if err != nil {
		http.Error(res, "Error parsing the file ", http.StatusInternalServerError)
		return
	}
	if name == "ascii-art.html" {

		err = tmpl.Execute(res, data)

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

		parseAndExecute(fileName, res, asciiS{})

	case "/ascii-art":
		if req.Method == http.MethodPost {
			fmt.Println("Post Post Post ")

		}
		fileName := "ascii-art.html"
		var a asciiS

		a.input = req.FormValue("input")
		a.banner = req.FormValue("banner")

		parseAndExecute(fileName, res, a)
	default:
		fileName := "404.html"
		http.Error(res, "Error parsing the file ", http.StatusInternalServerError)
		res.WriteHeader(404)
		parseAndExecute(fileName, res, asciiS{})
	}
}

func main() {
	const PORT = "8001"

	fs := http.FileServer(http.Dir("styles/"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))

	http.HandleFunc("/", handleFunc)

	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
