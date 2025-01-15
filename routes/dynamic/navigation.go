package dynamic

import "net/http"

type NavigationLink struct {
	Text string `json:"name"`
	URL  string `json:"url"`
}

var Router *http.ServeMux = http.NewServeMux()

func init() {

}
