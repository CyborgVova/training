package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var reg = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, s string)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		success := reg.FindStringSubmatch(r.URL.Path)
		if success == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, success[2])
	}
}

var tmpl = template.Must(template.ParseFiles("./tmpl/edit.html", "./tmpl/view.html"))

func renderTemplate(w http.ResponseWriter, page *Page, path string) {
	err := tmpl.ExecuteTemplate(w, path, page)
	if err != nil {
		error500(w, err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, page, "view.html")
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	err := page.save()
	if err != nil {
		error500(w, err)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, page, "edit.html")
}

func error500(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "%d %s", http.StatusInternalServerError, http.StatusText(500))
	log.Println(err)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))

	fmt.Println("Server is listen on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
