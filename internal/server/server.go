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

func CreateServer(logger *log.Logger) *HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("received %v request \"%v\" from \"%v\" (User-Agent: %v)", r.Method, r.URL, r.Host, r.UserAgent())
		if r.URL.Path != "/" {
			http.Error(w, "Not found", http.StatusNotFound)
			logger.Printf("client ERROR: %v invalid request parameter %v", http.StatusNotFound, r.URL)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			logger.Printf("client ERROR: %v method %v not allowed", http.StatusMethodNotAllowed, r.Method)
			return
		}
		handlers.HandleMain(logger)(w, r)
	})
	mux.HandleFunc("/upload", handlers.HandleUpload(logger))
	return &HTTPServer{
		Log: logger,
		Server: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},
	}
}
