package models

type Course struct {
	ID uint `json:"id" gorm:"primayKey"`
	Code string `json:"code"`
	Title string `json:"title"`
	Credit float32 `json:"credit"`

	Results []Result `json:"results" gorm:"foreignKey:CourseID"`
}