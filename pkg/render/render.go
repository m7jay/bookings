package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/m7jay/bookings/pkg/config"
	"github.com/m7jay/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds the data that should be available in all templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// Template renders template using html templates
func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {

		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	log.Println(tmpl)
	log.Println(tc)
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
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

// CreateTemplateCache looks in the templates folder and create a map of name to templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	// get all of the pages names *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}
	return cache, nil
}
