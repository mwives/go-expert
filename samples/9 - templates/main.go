package main

import (
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	// Most basic template
	course := Course{"Go", 40}

	tmpl := template.New("CourseTemplate")
	tmpl, err := tmpl.Parse("Course: {{.Name}} - Workload: {{.Workload}} hours")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}

	// Serving HTML template with partials
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// Create a new template and add custom functions to it
	tmpl = template.New("mainTemplate").Funcs(template.FuncMap{
		"upper": ToUpper,
	})

	// Parse the template files
	tmpl, err = tmpl.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "content.html", Courses{
			{"Go", 40},
			{"Python", 60},
			{"Java", 80},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
