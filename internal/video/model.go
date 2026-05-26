package video

import (
    "time"
    "github.com/google/uuid"
)


type Video struct {
    ID uuid.UUID `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    DurationSeconds int32 `json:"duration_seconds"`
    CreatedAt time.Time `json:"created_at"`
}
