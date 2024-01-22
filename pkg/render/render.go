package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/toastsandwich/bookings/pkg/config"
	"github.com/toastsandwich/bookings/pkg/models"
)

// SIMPLE RENDER FUNCTION
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.gohtml")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// map[string]*template.Template -> cache to avoid using hardware
// var templateCache = make(map[string]*template.Template)

// func RenderTemplateWithCache(w http.ResponseWriter, t string) {
// 	if _, ok := templateCache[t]; !ok {
// 		fmt.Println("creating template adding to cache")
// 		if err := createTemplateCache(t); err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("using cached template")
// 	}
// 	tmpl := templateCache[t]
// 	if err := tmpl.Execute(w, nil); err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{"./templates/" + t, "./templates/base.gohtml"}
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	templateCache[t] = tmpl
// 	return nil
// }

var app *config.AppConfig

// newTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	// get them template cache from AppConfig
	var cache = map[string]*template.Template{}
	if app.UseCache {
		cache = app.TemplateCache
	} else {
		cache, _ = CreateTemplateCache()
	}
	// // create a template cache
	// cache, err := CreateTemplateCache()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// get requested template form cache
	template, ok := cache[tmpl]
	if !ok {
		fmt.Println(tmpl, " not found")
	}
	data = addDefaultData(data)
	buf := new(bytes.Buffer)
	err := template.Execute(buf, data)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all files --> *.gohtml
	pages, err := filepath.Glob("./templates/*.gohtml")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			fmt.Println("error in parsing ", page)
		}
		ts, err = ts.ParseGlob("./templates/base.gohtml")
		if err != nil {
			return myCache, fmt.Errorf("error in ts.ParseGlob()")
		}
		myCache[name] = ts
	}
	return myCache, nil
}
