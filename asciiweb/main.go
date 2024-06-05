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
		if req.Method != http.MethodGet {
			res.WriteHeader(http.StatusBadRequest)

			return
		}
		fileName := "index.html"

		parseAndExecute(fileName, res, asciiS{})

	case "/ascii-art":
		if req.Method != http.MethodPost {

			fileName := "400.html"
			res.WriteHeader(http.StatusBadRequest)
			parseAndExecute(fileName, res, asciiS{})
			return
		}
		fileName := "ascii-art.html"
		var a asciiS

		a.input = req.FormValue("input")
		a.banner = req.FormValue("banner")

		parseAndExecute(fileName, res, a)
	default:
		fileName := "404.html"
		res.WriteHeader(http.StatusNotFound)
		parseAndExecute(fileName, res, asciiS{})
	}
}

func main() {
	const PORT = "8000"
	// create Dir of type http.Dir
	// dir := http.Dir("/static/")
	// http.Handle("/static/", http.FileServer(dir))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer("styles")))
	fs := http.FileServer(http.Dir("./static/styles"))
	http.Handle("/static/styles/", http.StripPrefix("/static/styles", fs))

	imgs := http.FileServer(http.Dir("./static/imgs"))
	http.Handle("/static/imgs/", http.StripPrefix("/static/imgs", imgs))

	http.HandleFunc("/", handleFunc)

	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
