package models 

type Result struct {
	ID uint `json:"id" gorm:"primarykey"`
	StudentID uint `json:"student_id" gorm:"not null"`
	Subject string `json:"subject"`
	Score int `json:"score"`
	Student Student `json:"student" gorm:"foreignKey:StudentID"`
}