package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID              uint    `json:"student_id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	Email           string    `json:"email" gorm:"unique"`
	Password        string    `json:"password"`

	Phone           string    `json:"phone"`
	DateOfBirth     string    `json:"date_of_birth"`
	Gender          string    `json:"gender"`
	BloodGroup      string    `json:"blood_group"`
	Address         string    `json:"address"`

	Department      string    `json:"department"`
	Program         string    `json:"program"`
	Batch           string    `json:"batch"`
	CurrentSemester string    `json:"current_semester"`
	CreditCompleted int       `json:"credit_completed"`
	TotalCredits    int       `json:"total_credits"`
	CGPA            float32   `json:"cgpa"`
	AdmissionDate   string    `json:"admission_date"`

	GuardianName     string  `json:"guardian_name"`
	GuardianPhone    string  `json:"guardian_phone"`
	GuardianRelation string  `json:"guardian_relation"`

	TotalPayable   float64   `json:"total_payable"`
	TotalPaid      float64   `json:"total_paid"`
	TotalDue       float64   `json:"total_due"`
	PaymentStatus  string    `json:"payment_status"`

	Results        []Result  `json:"results" gorm:"foreignKey:StudentID"`
}
