package models

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:VARCHAR(32 )"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:VARCHAR(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130" gorm:"type:INTEGER"`
	Email     string `json:"email" binding:"required,email" gorm:"type:VARCHAR(256)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=100" validate:"is-cool" gorm:"type:VARCHAR(200)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:VARCHAR(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:VARCHAR(200);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignKey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}
