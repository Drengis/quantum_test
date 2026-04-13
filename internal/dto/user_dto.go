package dto

type CreateUserRequest struct {
	TgID      string  `json:"tg_id"`
	Username  *string `json:"username"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	LangCode  *string `json:"lang_code"`
	InvitedBy *string `json:"invited_by"`
}
