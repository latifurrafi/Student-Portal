package models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	StudentID uint
	Course    string
	Grade     string
	Credit    int
	Semester  string
}
