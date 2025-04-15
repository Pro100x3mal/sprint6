package handlers

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	htmlFile   = "index.html"
	uploadFile = "myFile"
	uploadDir  = "uploads"
	maxSize    = 10 << 20
)

func HandleMain(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("received %v request \"%v\" from \"%v\" (User-Agent: %v)", r.Method, r.URL, r.Host, r.UserAgent())

		http.ServeFile(w, r, htmlFile)

		logger.Printf("the requested resource %v was successfully sent", htmlFile)
	}
}
func HandleUpload(sLog *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sLog.Printf("received %v request \"%v\" from \"%v\" (User-Agent: %v)", r.Method, r.URL, r.Host, r.UserAgent())

		err := r.ParseMultipartForm(maxSize)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}

		file, fileHeader, err := r.FormFile(uploadFile)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}
		defer file.Close()

		if fileHeader.Size > maxSize {
			http.Error(w, "File too large (10MB)", http.StatusRequestEntityTooLarge)
			sLog.Printf("client ERROR: %v the uploaded file \"%v\" exceeds the allowed size (%v > %v).", http.StatusRequestEntityTooLarge, fileHeader.Filename, fileHeader.Size, maxSize)
			return
		}

		err = os.Mkdir(uploadDir, 0755)
		if err != nil && !errors.Is(err, os.ErrExist) {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}

		outName := time.Now().UTC().String() + ".txt"
		outPath := filepath.Join(uploadDir, outName)
		out, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}
		defer out.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}

		result, err := service.ConvertString(string(data))
		if err != nil {
			http.Error(w, fmt.Sprintf(err.Error()), http.StatusInternalServerError)
			sLog.Println(err)
			return
		}

		_, err = out.WriteString(result + "\n")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}

		root, err := os.OpenRoot(uploadDir)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}
		defer root.Close()

		rootFile, err := root.Open(outName)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}
		defer rootFile.Close()

		info, err := rootFile.Stat()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
			return
		}
		w.Header().Set("Content-Disposition", "attachment; filename="+outName)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprint(info.Size()))
		_, err = io.Copy(w, rootFile)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sLog.Println(err)
		}

		sLog.Printf("the requested resource %v was successfully sent (%v bytes)", outName, info.Size())
	}
}
