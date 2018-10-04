package yext

import (
	"encoding/json"
)

const ENTITYTYPE_EVENT EntityType = "EVENT"

// TODO: "Event" conflicts with the Event struct in list.go, but should consider better name
type EventEntity struct {
	EntityMeta  *EntityMeta `json:"meta,omitempty"`
	Id          *string     `json:"id,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	EntityType  EntityType  `json:"entityType,omitempty"`
}

func (e *EventEntity) GetEntityId() string {
	if e.Id != nil {
		return *e.Id
	}
	return ""
}

func (e *EventEntity) GetEntityType() EntityType {
	return ENTITYTYPE_EVENT
}

func (e *EventEntity) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
