package models 

type Result struct {
	ID uint `json:"id" gorm:"primarykey"`
	StudentID uint `json:"student_id"`
	Subject string `json:"subject"`
	Score int `json:"score"`
}