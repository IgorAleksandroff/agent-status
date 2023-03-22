package grpchandlers

import (
	"fmt"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
)

var grpcToStatus = []entity.Status{
	entity.StatusActive,
	entity.StatusReqInactive,
	entity.StatusInactive,
	entity.StatusReqBreak,
	entity.StatusBreak,
	entity.StatusForceMajeure,
	entity.StatusChat,
	entity.StatusLetter,
}

func GetStatusIDbyName(name string) (int32, error) {
	for id, s := range grpcToStatus {
		if string(s) == name {
			return int32(id), nil
		}
	}

	return -1, fmt.Errorf("id not found for status - %s", name)
}
