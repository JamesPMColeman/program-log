package main

import (
	"log"
	"github.com/JamesPMColeman/program-log/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
