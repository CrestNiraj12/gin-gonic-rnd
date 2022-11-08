package services

import . "golab-gin-poc/models"

type VideoService interface {
	Save(Video) Video
	FindAll() []Video
}

type videoService struct {
	videos []Video
}

func New() VideoService {
	return &videoService{
		videos: make([]Video, 0),
	}
}

func (service *videoService) Save(video Video) Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []Video {
	return service.videos
}
