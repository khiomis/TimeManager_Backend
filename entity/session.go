package entity

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id        uuid.UUID `db:"id_session"`
	CreatedAt time.Time `json:"created_at" db:"dt_created_at"`
	IdUser    int64     `json:"id_user" db:"id_user"`
	ExpireAt  time.Time `json:"expire_at" db:"dt_expires_at"`
}
