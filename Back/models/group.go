package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Users []User `json:"users"`
}
