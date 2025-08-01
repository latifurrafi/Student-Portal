package models

type Student struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Results []Result `json:"results" gorm:"foreignKey:StudentID"`
}