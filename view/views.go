package view

import (
	. "glu-software/internal"
	"html/template"
)

var Templates *template.Template

func init() {

	// Initialized data
	var err error = nil

	// Parse go templates
	Templates, err = template.ParseGlob("view/*.tmpl")

	// Error check
	CheckErr(err)
}
