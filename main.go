package main

import (
	"fmt"
	"log"
	"net/http"
	"videoserver/app/controllers"
	"videoserver/app/controllers/views"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", views.IndexView)

	r.Handle("/static/{dir}/{file}", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/api/video/{video_id}/{filename}", controllers.RequestHandler)

	corsObj := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Started server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(r)))

}
