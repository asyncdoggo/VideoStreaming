package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"videoserver/app/controllers/models"

	"github.com/google/uuid"
)

type FilesResponse struct {
	Status string       `json:"status"`
	Files  []FileObject `json:"files"`
}

type FileUploadResponse struct {
	Status string `json:"status"`
}

type FileObject struct {
	Filename string `json:"name"`
	FileId   string `json:video_id`
	Filesize int64  `json:"size"`
}

func ServeVideo(w http.ResponseWriter, r *http.Request, video_id string, filename string) {
	http.ServeFile(w, r, fmt.Sprintf("app/videos/%s/%s", video_id, filename))
}

func UploadVideo(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(104857600) // 100 mb limit
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	uid := uuid.New()

	err = os.Mkdir(fmt.Sprintf("app/videos/%s", uid.String()), fs.FileMode(os.O_CREATE))

	models.Db.Create(&models.Files{FileID: uid.String(), Filename: header.Filename, FileSize: header.Size})

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	f, err := os.OpenFile(fmt.Sprintf("app/videos/%s/video.mp4", uid.String()), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	go process_video(uid, w, r)

	resp := FileUploadResponse{Status: "success"}

	response, err := json.Marshal(resp)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func process_video(uid uuid.UUID, w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	absPath, _ := filepath.Abs(fmt.Sprintf("app/videos/%s", uid.String()))

	cmd := exec.Command("ffmpeg", "-i", fmt.Sprintf("%s\\video.mp4", absPath), "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", fmt.Sprintf("%s\\video.m3u8", absPath))
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	err := os.Remove(fmt.Sprintf("%s\\video.mp4", absPath))

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprint(w, "success")

}

func VideoList(w http.ResponseWriter, r *http.Request) {
	rows, err := models.Db.Model(&models.Files{}).Rows()

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	defer rows.Close()

	farr := []FileObject{}

	resp := FilesResponse{"success", farr}
	for rows.Next() {
		var file models.Files
		models.Db.ScanRows(rows, &file)
		f := FileObject{file.Filename, file.FileID, file.FileSize}
		farr = append(farr, f)
	}

	resp.Files = farr

	response, err := json.Marshal(resp)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
