package controllers

import (
	. "golab-gin-poc/models"
	"golab-gin-poc/services"
	"golab-gin-poc/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service services.VideoService
}

var validate *validator.Validate

func New(service services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

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

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}
