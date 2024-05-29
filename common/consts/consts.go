package consts

import "gorm.io/gorm"

var ERROR_NOT_FOUND = gorm.ErrRecordNotFound

const (
	GENDER_MAN = iota + 1
	GENDER_WOMAN
)

type MessageResponse struct {
	Model     string        `json:"model"`
	CreatedAt string        `json:"created_at"`
	Message   MessageEntity `json:"message"`
	Done      bool          `json:"done"`
}

type MessageEntity struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

const ASSISTANT = "assistant"
