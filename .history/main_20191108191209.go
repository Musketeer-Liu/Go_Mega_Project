package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// User struct
type User struct {
	Username string
}

// Post struct
type Post struct {
	User // 这里可以采用匿名简化 不用写User User HTML模板中也可以不用再写.User了
	// User User
	Body string
}

// IndexViewModel struct
type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

// PopulateTemplates func
// Create map template name to template.Template
// Unlike Flask, we need to hand write a function to extend from base template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}
			defer f.Close()
			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}
			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}
			result[fi.Name()] = tmpl
		}()
	}
	return result
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{Username: "Musketeer"}
		u2 := User{Username: "Paladin"}

		posts := []Post{
			Post{User: u1, Body: "Beautiful day in Portland!"},
			Post{User: u2, Body: "The Avengers movie was so cool!"},
		}

		v := IndexViewModel{Title: "Homepage", User: u1, Posts: posts}
		templates := PopulateTemplates()
		templates["index.html"].Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
