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

    videoService := video.NewService(database.New(db))
    videoHandler := video.NewHandler(videoService)

    mux := http.NewServeMux()

    mux.HandleFunc("POST /videos", videoHandler.CreateVideoHandler)
    mux.HandleFunc("GET /videos/{videoID}", videoHandler.GetVideoByIDHandler)
    mux.HandleFunc("GET /videos", videoHandler.GetVideosHandler)

    server := &http.Server{
        Handler: mux,
        Addr: ":8080",
    }


    fmt.Println("Starting server...")
    server.ListenAndServe()
}
