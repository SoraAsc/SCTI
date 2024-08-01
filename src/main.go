package main

import (
  "SCTI/fileserver"
  "SCTI/middleware"
  "SCTI/database"
  "log"
  "net/http"
  "github.com/joho/godotenv"
  "github.com/lengzuo/supa"
  "os"
  "fmt"
)

func start() (*supabase.Client) {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }



  conf := supabase.Config{
    ApiKey:     os.Getenv("SUPABASE_KEY"), 
    ProjectRef: os.Getenv("SUPABASE_URL"),
    Debug:      true,
  }
  supaClient, err := supabase.New(conf)
  if err != nil {
      fmt.Printf("failed in initialise client with err: %s", err)
      panic("FUCK")
  }
  return supaClient
}

type SignUpRequest struct {
  Email string
  Password string
}

func main() {
  err := database.OpenDatabase()
  if err != nil {
    log.Printf("Error connecting to postgres database\n%v", err)
  }
  defer database.CloseDatabase()

  supaClient := start()
  fileserver.RunFileServer()

  fmt.Println("Server Started")
  mux := http.NewServeMux()
  LoadRoutes(mux, supaClient)

  server := http.Server{
    Addr:    ":8080",
    Handler: middleware.EndpointLogging(mux),
  }

  log.Fatal(server.ListenAndServe())
}
