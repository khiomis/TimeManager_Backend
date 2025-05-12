package entity

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id        uuid.UUID `db:"id_session"`
	CreatedAt time.Time `json:"createdAt" db:"dt_created_at"`
	IdUser    int64     `json:"idUser" db:"id_user"`
	ExpireAt  time.Time `json:"expireAt" db:"dt_expires_at"`
}
