package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model

	Role        string
	Company     string
	Location    string
	Description string
	Salary      int64
	Remote      bool
	Link        string
}
