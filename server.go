package main

import (
	"guisteglich/WebService-With-Gin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", listAllHandler)
	server.POST("/videos", saveHandler)

	server.Run(":5050")
}

func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"massage": "ok",
	})

}

func (c *gin.Context)listAllHandler{
	c.JSON(200, videoController.FindAll())
}

func (c *gin.Context)saveHandler{
	c.JSON(200, videoController.Save(c)
}

