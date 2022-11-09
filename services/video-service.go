package services

import (
	"golab-gin-poc/models"
	"golab-gin-poc/repositories"
)

type VideoService interface {
	Save(models.Video) models.Video
	Update(models.Video) models.Video
	Delete(models.Video) models.Video
	FindAll() []models.Video
}

type videoService struct {
	videoRepository repositories.VideoRepository
}

func New(repo repositories.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video models.Video) models.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video models.Video) models.Video {
	service.videoRepository.Update(video)
	return video
}

func (service *videoService) Delete(video models.Video) models.Video {
	service.videoRepository.Delete(video)
	return video
}

func (service *videoService) FindAll() []models.Video {
	return service.videoRepository.FindAll()
}
