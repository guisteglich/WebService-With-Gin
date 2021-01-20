package service

import (
	"guisteglich/WebService-With-Gin/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(newVideo entity.Video) entity.Video {
	service.videos = append(service.videos, newVideo)
	return newVideo
}
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
