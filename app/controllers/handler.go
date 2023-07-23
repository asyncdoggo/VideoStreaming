package controllers

import (
	"net/http"
	"videoserver/app/controllers/api"

	"github.com/gorilla/mux"
)

func ServeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		api.ServeVideo(w, r, vars["video_id"], vars["filename"])
	}
}

func UploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		api.UploadVideo(w, r)
	}
}

func VideoListHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		api.VideoList(w, r)
	}
}
