package schemas

import (
	"time"

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

type OpeningResponse struct {
	ID          uint      `json:"id"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Salary      int64     `json:"salary"`
	Remote      *bool     `json:"remote"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
