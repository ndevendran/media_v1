package main

import (
    "fmt"
    "os"
    "net/http"
    "github.com/ndevendran/media_v1/internal/database"
    "github.com/joho/godotenv"
    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/ndevendran/media_v1/internal/video"
    "database/sql"
)

type apiConfig struct {
    db *database.Queries
}


func main() {
    err := godotenv.Load()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    dbURL := os.Getenv("DB_URL")
    db, err := sql.Open("pgx", dbURL)

    if err != nil {
    	fmt.Println(err)
    	os.Exit(1)
    }

    videoHandler := &video.Handler{}
    videoHandler.DB = database.New(db)

    mux := http.NewServeMux()

    mux.HandleFunc("POST /videos", videoHandler.CreateVideoHandler)
    mux.HandleFunc("POST /videos/{videoID}", videoHandler.GetVideoByIDHandler)

    server := &http.Server{
        Handler: mux,
        Addr: ":8080",
    }


    fmt.Println("Starting server...")
    server.ListenAndServe()
}
