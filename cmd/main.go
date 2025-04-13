package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	sLog := log.New(logFile, "", log.LstdFlags)

	s := server.CreateServer(sLog)
	s.Log.Printf("starting server on port %v", s.Server.Addr)
	err = s.Server.ListenAndServe()
	if err != nil {
		s.Log.Fatal(err)
	}
}
