package models

import "gorm.io/gorm"

type Semester struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Code int `json:"code"`

	Result []Result `json:"results" gorm:"foreignKey:SemesterID"`
}
