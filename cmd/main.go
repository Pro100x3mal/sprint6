package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	s := server.CreateServer(logger)
	logger.Printf("starting server on port %v", s.Server.Addr)
	err := s.Server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
