package video

import (
    "net/http"
    "github.com/ndevendran/media_v1/internal/database"
    "encoding/json"
    "context"
    "fmt"
    "github.com/google/uuid"
)

type Handler struct {
    DB *database.Queries
}


func (h *Handler) CreateVideoHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    video := Video{}
    err := decoder.Decode(&video)
    if err != nil {
        errorMsg := fmt.Sprintf("Internal Server Error: %v", err)
        http.Error(w, errorMsg, http.StatusInternalServerError)
    }

    createdVideo, err := h.DB.CreateVideo(context.Background(), database.CreateVideoParams{
        Title: video.Title,
        Description: video.Description,
        DurationSeconds: video.DurationSeconds,
    })

    if err != nil {
        errorMsg := fmt.Sprintf("Internal Server Error: %v", err)
        http.Error(w, errorMsg, http.StatusInternalServerError)
    }

    data, _ := json.Marshal(Video{
        ID: createdVideo.ID,
        Title: createdVideo.Title,
        Description: createdVideo.Description,
        CreatedAt: createdVideo.CreatedAt,
        DurationSeconds: createdVideo.DurationSeconds,
    })

    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(data)
	return
}

func (h *Handler) GetVideoByIDHandler(w http.ResponseWriter, r *http.Request) {
    	pathID := r.PathValue("videoID")

	    videoID, err := uuid.Parse(pathID)

        if err != nil {
            errorMsg := fmt.Sprintf("Internal Server Error: %v", err)
            http.Error(w, errorMsg, http.StatusBadRequest)
        }

        video, err := h.DB.GetVideoByID(context.Background(), videoID)

        if err != nil {
            errorMsg := fmt.Sprintf("Internal Server Error: %v", err)
            http.Error(w, errorMsg, http.StatusInternalServerError)
        }

        data, _ := json.Marshal(Video{
            ID: video.ID,
            Title: video.Title,
            Description: video.Description,
            CreatedAt: video.CreatedAt,
            DurationSeconds: video.DurationSeconds,
        })

        w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(200)
    	w.Write(data)
    	return
}
