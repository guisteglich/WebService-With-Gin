package controller

import (
	"guisteglich/WebService-With-Gin/entity"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) entity.Video
}

type controller struct {
	service service.videoService
}

type New(service service.VideoService) VideoController{
	return controller{
		service: service,
	}
}

func (control *controller)FindAll() []entity.Video{
	return control.service.FindAll()
}
func (control *controller)Save(c *gin.Context){
	var video entity.Video
	c.BindJSON(&video)
	c.service.Save(video)

	return video
}
