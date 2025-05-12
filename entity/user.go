package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id                int64      `json:"id" db:"id_user"`
	Uuid              uuid.UUID  `json:"uuid" db:"id_user_uuid"`
	CreatedAt         time.Time  `json:"createdAt" db:"dt_created_at"`
	UpdatedAt         time.Time  `json:"updatedAt" db:"dt_updated_at"`
	Name              string     `json:"name" db:"nm_user"`
	Email             string     `json:"email" db:"ds_email"`
	Password          string     `json:"password" db:"ds_password"`
	Status            UserStatus `json:"status" db:"tp_status"`
	TemporaryPassword string     `json:"temporaryPassword" db:"ds_temporary_password"`
}

type UserStatus int

const (
	UserPending UserStatus = iota
	UserActivated
	UserDeactivated
	UserBanned
)
