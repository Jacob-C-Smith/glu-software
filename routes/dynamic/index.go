package dynamic

import (
	. "glu-software/internal"
	"glu-software/routes"
	"glu-software/view"
	"net/http"
)

func init() {
	routes.Router.HandleFunc("/", HandleFunc)
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	err := view.Templates.ExecuteTemplate(w, "index.tmpl", "Hello")
	CheckErr(err)
}
