package controllers

import (
	. "golab-gin-poc/models"
	"golab-gin-poc/services"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
