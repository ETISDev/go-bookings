package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate parses and executes a template file using text/template
func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.gotpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// TemplateCache is a map that stores parsed templates
var templateCache = make(map[string]*template.Template)

func RenderTemplateWithTemplateCache(w http.ResponseWriter, templateName string) {
	var err error
	var tmpl *template.Template
	var ok bool

	// Check if the template is already in the cache
	_, ok = templateCache[templateName]
	if !ok {
		// If not, parse the template and add it to the cache
		tmpl, err = template.ParseFiles("./templates/"+templateName, "./templates/base.layout.gotpl")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templateCache[templateName] = tmpl
	}

	tmpl = templateCache[templateName]
	// Execute the template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

// CreateTemplateCache parses and caches templates
func CreateTemplateCache() (templateName string) error {
	// Create a new map to store the template cache
	// templateCache := make(map[string]*template.Template)

	// List of template files to parse
	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.gotpl"	,
	}

	// Loop through the template files and parse them
	// for _, templateName := range templates {
		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			return nil, err
		}
		templateCache[templateName] = tmpl
	// }

	return  nil
}

