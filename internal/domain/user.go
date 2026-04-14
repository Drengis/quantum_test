package domain

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TgID      string    `json:"tg_id" gorm:"uniqueIndex;not null"`
	Username  *string   `json:"username"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  *string   `json:"last_name"`
	LangCode  *string   `json:"lang_code"`
	InvitedBy *string   `json:"invited_by"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
