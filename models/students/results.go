package models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`

	StudentID uint `json:"student_id"`
	CourseID uint `json:"course_id"`
	SemesterID uint `json:"semester_id"`

	Grade string `json:"grade"`
	GradePoint float32 `json:"grade_point"`

	Student Student `json:"student" gorm:"foreignKey:StudentID"`
	Course Course `json:"course" gorm:"foreignKey:CourseID"`
	Semester Semester `json:"semester" gorm:"foreignKey:SemesterID"`
}