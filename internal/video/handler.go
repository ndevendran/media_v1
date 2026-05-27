package video

import (
    "net/http"
    "encoding/json"
    "log"
    "strconv"
)

type Handler struct {
    Service *Service
}

func NewHandler(service *Service) *Handler {
    return &Handler{
        Service: service,
    }
}

func (h *Handler) CreateVideoHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    request := CreateVideoRequest{}
    err := decoder.Decode(&request)
    if err != nil {
        log.Printf("Internal Server Error: %v", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    video, err := h.Service.CreateVideo(r.Context(), request)

    if err != nil {
        switch err.Error() {
            case "title is required",
                "title is too long",
                "duration must be greater than zero":
                    http.Error(
                        w,
                        err.Error(),
                        http.StatusBadRequest,
                    )
            default:
                log.Printf("create video: %v", err)
                http.Error(
                    w,
                    "internal server error",
                    http.StatusInternalServerError,
                )
        }

        return
    }

    response, err := json.Marshal(Video{
        ID: video.ID,
        Title: video.Title,
        Description: video.Description,
        CreatedAt: video.CreatedAt,
        DurationSeconds: video.DurationSeconds,
    })

    if err != nil {
        log.Printf("Internal Server Error: %v", err)
        http.Error(
            w,
            "internal server error",
            http.StatusInternalServerError,
        )
        return
    }

    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (h *Handler) GetVideoByIDHandler(w http.ResponseWriter, r *http.Request) {
    	pathID := r.PathValue("videoID")

        video, err := h.Service.GetVideoByID(r.Context(), pathID)

        if err != nil {
            switch err.Error() {
                case "could not parse video id":
                        http.Error(
                            w,
                            err.Error(),
                            http.StatusBadRequest,
                        )
                case "video not found":
                    http.Error(
                        w,
                        err.Error(),
                        http.StatusNotFound,
                    )
                default:
                    log.Printf("create video: %v", err)
                    http.Error(
                        w,
                        "internal server error",
                        http.StatusInternalServerError,
                    )
            }

            return
        }

        data, err := json.Marshal(Video{
            ID: video.ID,
            Title: video.Title,
            Description: video.Description,
            CreatedAt: video.CreatedAt,
            DurationSeconds: video.DurationSeconds,
        })

        if err != nil {
            log.Printf("Internal Server Error: %v", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(200)
    	w.Write(data)
}

func (h *Handler) GetVideosHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    limit, err := strconv.Atoi(query.Get("limit"))

    if err != nil {
        http.Error(w, "could not parse limit", http.StatusBadRequest)
        return
    }

    offset, err := strconv.Atoi(query.Get("offset"))

    if err != nil {
        http.Error(w, "could not parse offset", http.StatusBadRequest)
        return
    }

    req := GetVideosRequest{
        Limit: int32(limit),
        Offset: int32(offset),
        OrderBy: query.Get("sort"),
    }

    videos, err := h.Service.GetVideos(r.Context(), req)

    if err != nil {
        switch err.Error() {
        case "offset must be greater than 0":
                    http.Error(
                        w,
                        err.Error(),
                        http.StatusBadRequest,
                    )
            case "invalid sort option":
                http.Error(
                    w,
                    err.Error(),
                    http.StatusNotFound,
                )
            default:
                log.Printf("Get Videos Error: %v", err)
                http.Error(
                    w,
                    "internal server error",
                    http.StatusInternalServerError,
                )
        }

        return
    }

    data, err := json.Marshal(videos)

    if err != nil {
        log.Printf("Internal Server Error: %v", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(data)
}
