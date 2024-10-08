package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	NbStudents int `json:"nbStudents"`
	NbGroups int `json:"nbGroups"`
	LastMin int `json:"lastMin"`
	LastMax int `json:"lastMax"`
}
