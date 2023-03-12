package entity

type CommandType string

const (
	SendMsg        CommandType = "send message"
	autoAssignment CommandType = "update auto-assignment"
)

type Event struct {
	Type   CommandType
	Params *map[string]string
}
