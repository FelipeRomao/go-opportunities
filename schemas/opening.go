package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint   `json:"id"`
	Role      string `json:"role"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	Remote    bool   `json:"remote"`
	Link      string `json:"link"`
	Salary    int64  `json:"salary"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}
