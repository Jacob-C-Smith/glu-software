// Package
package routes

// Imports
import (
	. "glu-software/internal"
	"glu-software/internal/colorful"
	"glu-software/routes/static"
	"glu-software/view"
	"net/http"
)

// Data
var Router *http.ServeMux = nil

// Early
func init() {

	// Construct a new multiplexer
	Router = http.NewServeMux()

	// Set up the static content router
	Router.Handle("/static/", http.StripPrefix("/static", static.Router))

	// Set up the dynamic content router(s)
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Initialized data
		var err error = nil

		// Route requests
		if r.URL.String() == "/" {

			// Index template
			view.Templates.ExecuteTemplate(w, "index.tmpl", "Glu Software")
		} else if r.URL.String() == "/about" {

			// About template
			view.Templates.ExecuteTemplate(w, "about.tmpl", "Glu Software")
		} else {

			// Navigation bar template
			view.Templates.ExecuteTemplate(w, "navbar.tmpl", nil)
		}

		// Error checking
		CheckErr(err)
	})

	// Logs
	colorful.Log(colorful.LogLevels.Info, "[glu] [router] All routes registered\n")
}
