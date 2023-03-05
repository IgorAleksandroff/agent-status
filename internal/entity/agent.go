package entity

import "time"

type Agent struct {
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"password" db:"password"`
	StatusID  int64     `db:"status_id"`
	CreatedAt time.Time `db:"created_at"`
}
