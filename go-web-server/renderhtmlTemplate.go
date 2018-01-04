package main

import (
	"html/template"
	"net/http"
	"path"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", bar)
	http.ListenAndServe(":3000", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Yang", []string{"play", "programming"}}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// buf := new(bytes.Buffer)
	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// templateString := buf.String()
	// fmt.Println(templateString, "...")
}
