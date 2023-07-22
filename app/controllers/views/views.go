package views

import "net/http"

func IndexView(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "app/views/templates/index.html")
}
