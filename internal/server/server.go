package server

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"log"
	"net/http"
	"time"
)

const (
	addr         = ":8080"
	readTimeout  = time.Second * 5
	writeTimeout = time.Second * 10
	idleTimeout  = time.Second * 15
)

type HTTPServer struct {
	Log    *log.Logger
	Server *http.Server
}

func CreateServer(sLog *log.Logger) *HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleMain(sLog))
	mux.HandleFunc("/upload", handlers.HandleUpload(sLog))
	return &HTTPServer{
		Log: sLog,
		Server: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ErrorLog:     sLog,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},
	}
}
