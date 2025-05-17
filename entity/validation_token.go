package entity

import (
	"github.com/google/uuid"
	"time"
)

type ValidationToken struct {
	Id        uuid.UUID           `json:"id" db:"id_validation_token"`
	CreatedAt time.Time           `json:"created_at" db:"dt_created_at"`
	IdUser    int64               `json:"id_user" db:"id_user"`
	Code      string              `json:"code" db:"cd_validation_token"`
	ExpireAt  time.Time           `json:"expire_at" db:"dt_expire_at"`
	Type      ValidationTokenType `json:"type" db:"tp_validation_token"`
}

type ValidationTokenType int

const (
	ValidationTokenTypeSignIn ValidationTokenType = iota
	ValidationTokenTypePasswordRecovery
	ValidationTokenTypeAccountUpdate
)
