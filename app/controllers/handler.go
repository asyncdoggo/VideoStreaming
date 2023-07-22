package controllers

import (
	"net/http"
	"videoserver/app/controllers/api"

	"github.com/gorilla/mux"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		api.ServeVideo(w, r, vars["video_id"], vars["filename"])
	}
}
