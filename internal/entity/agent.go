package entity

import "time"

type Agent struct {
	Login     string    `json:"login" db:"login"`
	Password  *string   `json:"password,omitempty" db:"password"`
	Status    *Status   `json:"status,omitempty" db:"status"`
	CreatedAt time.Time `db:"created_at"`
}
