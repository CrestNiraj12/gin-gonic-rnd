package repositories

import (
	"golab-gin-poc/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video models.Video)
	Update(video models.Video)
	Delete(video models.Video)
	FindAll() []models.Video
}

type database struct {
	connection *gorm.DB
}

func New() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Video{}, &models.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(video models.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video models.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video models.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []models.Video {
	var videos []models.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
