package yext

import (
	"encoding/json"
)

const ENTITYTYPE_EVENT EntityType = "event"

type EventEntity struct {
	BaseEntity
	Id          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	EntityType  EntityType `json:"entityType,omitempty"`
}

func (e *EventEntity) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
