package model

// System is system.
type System struct {
	Model
	Name    string `json:"name,omitempty" gorm:"size:50;not null" valid:"required,runelength(1|50)"`
	Company string `json:"company,omitempty" gorm:"size:50" valid:"runelength(1|50)"`
	Mail    string `json:"mail,omitempty" gorm:"size:255;not null" valid:"required,runelength(1|255)"`
}
