package yext

import (
	"encoding/json"

	"github.com/mohae/deepcopy"
)

const (
	ENTITYTYPE_EVENT     EntityType = "EVENT"
	EntityPathNameEvents            = "events" // TODO: rename
)

type EventEntity struct { // TODO: rename
	//EntityMeta
	Id          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	EntityType  EntityType `json:"entityType,omitempty"`
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

func (e *EventEntity) PathName() string {
	return EntityPathNameEvents
}

func (y *EventEntity) Copy() Entity {
	return deepcopy.Copy(y).(*EventEntity)
}

func (e *EventEntity) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
