package static

import (
	"glu-software/internal/colorful"
	"net/http"
)

var Router *http.ServeMux

func init() {

	// Construct a new multiplexer for static content
	Router = http.NewServeMux()

	//
	Router.Handle("/", http.FileServer(http.Dir("./static")))

	// Logs
	colorful.Log(colorful.LogLevels.Info, "[glu] [router] [static] Static router registered\n")
}
