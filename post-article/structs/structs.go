package structs

import (
	"time"
)

type Posts struct {
	Id           int       `json:"id" gorm:"primary_key;auto_increment"`
	Title        string    `json:"title" validate:"min=10,max=200"`
	Content      string    `json:"content" validate:"min=200"`
	Category     string    `json:"category" validate:"min=3,max=100"`
	Created_date time.Time `json:"created_date"`
	Updated_date time.Time `json:"updated_date"`
	Status       string    `json:"status" validate:"eq=Publish|eq=Draft|eq=Thrash"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
