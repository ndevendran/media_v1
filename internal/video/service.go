package video

import (
    "github.com/ndevendran/media_v1/internal/database"
    "context"
    "errors"
    "github.com/google/uuid"
    "fmt"
)

type Service struct {
    DB *database.Queries
}

func NewService(db *database.Queries) *Service {
    return &Service{
        DB: db,
    }
}


func (s *Service)  CreateVideo(
    ctx context.Context,
    req CreateVideoRequest,
)(database.Video, error) {
    // validation
    if req.Title == "" {
        return database.Video{}, errors.New("title is required")
    }

    if len(req.Title) > 200 {
        return database.Video{}, errors.New("title is too long")
    }

    if req.DurationSeconds <= 0 {
        return database.Video{}, errors.New("duration must be greater than 0")
    }

    video, err := s.DB.CreateVideo(ctx, database.CreateVideoParams{
        Title: req.Title,
        Description: req.Description,
        DurationSeconds: req.DurationSeconds,
    })

    if err != nil {
        return database.Video{}, err
    }

	return video, nil
}

func (s *Service) GetVideoByID(
    ctx context.Context,
    pathID string,
)(database.Video, error) {
    videoID, err := uuid.Parse(pathID)

    if err != nil {
        return database.Video{}, errors.New("could not parse video id")
    }

    video, err := s.DB.GetVideoByID(ctx, videoID)

    if err != nil {
        return database.Video{}, errors.New("video not found")
    }

    return video, nil
}

func (s *Service) GetVideos(
    ctx context.Context,
    req GetVideosRequest,
)([]database.Video, error) {
    if req.Offset < 0 {
        return []database.Video{}, fmt.Errorf("offset must be greater than 0")
    }

    if req.Limit > 100 {
        req.Limit = 100
    }

    if req.OrderBy == "asc" {
        videos, err := s.DB.GetVideosAsc(ctx, database.GetVideosAscParams{
            Limit: req.Limit,
            Offset: req.Offset,
        })

        if err != nil {
            return []database.Video{}, fmt.Errorf("internal server error: %v", err)
        }

        return videos, nil
    }

    if req.OrderBy == "desc" {
        videos, err := s.DB.GetVideosDesc(ctx, database.GetVideosDescParams{
            Limit: req.Limit,
            Offset: req.Offset,
        })

        if err != nil {
            return []database.Video{}, fmt.Errorf("internal server error: %v", err)
        }

        return videos, nil
    }

    if req.OrderBy == "" {
        videos, err := s.DB.GetVideos(ctx, database.GetVideosParams{
            Limit: req.Limit,
            Offset: req.Offset,
        })

        if err != nil {
            return []database.Video{}, fmt.Errorf("internal server error: %v", err)
        }

        return videos, nil
    }

    return []database.Video{}, fmt.Errorf("invalid sort option")
}
