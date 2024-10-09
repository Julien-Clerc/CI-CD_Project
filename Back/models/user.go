package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	GroupID uint `json:group_id`
	Group Group `json:"group"`
}
