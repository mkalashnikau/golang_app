package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// FuncMap is a map of functions that we can use in a template
var functions = template.FuncMap{}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache() // btw making cache is not sufficient approach
	if err != nil {
		log.Fatal(err) // It stops the app in case of error
	}

	// we are adding ok varible to check if the file exists or not
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// Buffer which will hold information
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil) // gives the template, doesn't store anything to the file and put the content to the buffer

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// CreateTemplateCache creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Define if we need to process the template or leave as is
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
