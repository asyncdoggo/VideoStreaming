package api

import (
	"fmt"
	"net/http"
)

func ServeVideo(w http.ResponseWriter, r *http.Request, video_id string, filename string) {
	http.ServeFile(w, r, fmt.Sprintf("app/videos/%s/%s", video_id, filename))
}
