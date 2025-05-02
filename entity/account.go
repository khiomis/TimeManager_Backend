package entity

import "time"

type Account struct {
	Id        string        `json:"id" sql:"type:varchar(40);primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time     `json:"createdAt" sql:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time     `json:"UpdatedAt" sql:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Email     string        `json:"email" sql:"type:varchar(128);unique_index;not null"`
	Name      string        `json:"name" sql:"type:varchar(128);unique_index;not null"`
	Password  string        `json:"password" sql:"type:varchar(128);unique_index;not null"`
	Status    AccountStatus `json:"status" sql:"type:integer"`
}

type AccountStatus int

const (
	AccountPending AccountStatus = iota
	AccountActivated
	AccountDeactivated
	AccountBanned
)
