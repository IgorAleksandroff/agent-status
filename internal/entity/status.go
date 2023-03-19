package entity

import "time"

type (
	Mode   string
	Status string
)

const (
	Rest Mode = "MAN"
	Grpc Mode = "AUT"
)

const (
	StatusActive       Status = "active"
	StatusReqInactive  Status = "request to inactive"
	StatusInactive     Status = "inactive"
	StatusReqBreak     Status = "request to break"
	StatusBreak        Status = "break"
	StatusForceMajeure Status = "force majeure"
	StatusChat         Status = "chat"
	StatusLetter       Status = "letter"
)

type Transition struct {
	Src  Status `json:"source" db:"source"`
	Dst  Status `json:"destination" db:"destination"`
	Mode Mode   `json:"mode" db:"mode"`
}

type Logs struct {
	Agent       string    `json:"agent" db:"agent_login"`
	ProcessedAt time.Time `json:"processed_at" db:"processed_at"`
	Transition
}
