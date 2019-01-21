package model

// System is system.
type System struct {
	Model
	Name    string `json:"name" gorm:"size:50;not null"`
	Company string `json:"company" gorm:"size:50"`
	Mail    string `json:"mail" gorm:"size:255;not null"`
}
