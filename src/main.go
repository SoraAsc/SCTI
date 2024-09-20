package main

import (
	"SCTI/database"
	"SCTI/fileserver"
	"SCTI/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.OpenDatabase()
	if err != nil {
		log.Printf("Error connecting to postgres database\n%v", err)
	}
	defer database.CloseDatabase()

	fileserver.RunFileServer()

	mux := http.NewServeMux()
	LoadRoutes(mux)

	server := http.Server{
    Addr:    ":8080", // adicione :xx na URL do .env se a porta não for a padrão do protocolo
		Handler: middleware.EndpointLogging(mux),
	}

  fmt.Printf("Server Started at: %s\n", os.Getenv("URL"))
	log.Fatal(server.ListenAndServe())
}
