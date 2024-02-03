package render

import (
	"MyFirstApp/pkg/config"
	"MyFirstApp/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData used for Globally Sharing template Data
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RrTemplate(w http.ResponseWriter, fileHtml string, td *models.TemplateData) {
	tCache := map[string]*template.Template{}

	//app.UseCache reading from disk value is False can be change in main.go
	if app.UseCache {
		tCache = app.TemplateCache
	} else {
		tCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tCache[fileHtml]
	if !ok {
		log.Fatal("coud not get templates from templates cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	// adding template and data to buffer
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//get all files from templates package named */.page.gohtml
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with  */.page.gohtml
	for _, page := range pages {
		//separation only name from path to the page
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.page.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.page.gohtml")
			if err != nil {
				return myCache, nil
			}

			myCache[name] = ts
		}
	}

	return myCache, err
}
