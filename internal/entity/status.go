package entity

type mode string

const (
	Rest mode = "MAN"
	Grpc mode = "AUT"
)

// StatusActive - начало смены
const StatusActive int64 = 1

type Transition struct {
	StatusID     int64   `json:"status_id" db:"status_id"`
	PermittedIDs []int64 `json:"permitted_ids" db:"permitted_ids"`
}

type Logs struct {
	Agent       string `json:"agent" db:"agent_login"`
	OldStatusID int64  `json:"old_status_id" db:"old_status_id"`
	NewStatusID int64  `json:"new_status_id" db:"new_status_id"`
	Mode        mode   `json:"mode" db:"mode"`
	ProcessedAt string `json:"processed_at" db:"processed_at"`
}
