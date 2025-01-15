package dynamic

import (
	. "glu-software/internal"
	"glu-software/routes"
	"glu-software/view"
	"net/http"
)

func init() {
	routes.Router.HandleFunc("/about", handleFunc)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	err := view.Templates.ExecuteTemplate(w, "about.tmpl", "Hello")
	CheckErr(err)
}
