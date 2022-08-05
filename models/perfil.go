package models

type Usuario struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Data  string `json:"lastName"`
}
