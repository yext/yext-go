package yext

type ActivityLog struct {
	TimeStamp    *int        `json:"timestamp"`
	LocationID   *string     `json:"locationId"`
	Details      *string     `json:"details"`
	Content      *string     `json:"content"`
	Type         *string     `json:"type"`
	Actor        *string     `json:"actor"`
	ActorDetails ActorDetail `json:"actorDetails"` // need to set version 20210728
}

type ActorDetail struct {
	Name  *string `json:"name"`
	Email *string `json:"emails"`
}

func (y ActivityLog) GetTimeStamp() int {
	if y.TimeStamp != nil {
		return *y.TimeStamp
	}
	return 0
}

func (y ActivityLog) GetLocationID() string {
	if y.LocationID != nil {
		return *y.LocationID
	}
	return ""
}

func (y ActivityLog) GetDetails() string {
	if y.Details != nil {
		return *y.Details
	}
	return ""
}

func (y ActivityLog) GetContent() string {
	if y.Content != nil {
		return *y.Content
	}
	return ""
}

func (y ActivityLog) GetType() string {
	if y.Type != nil {
		return *y.Type
	}
	return ""
}

func (y ActivityLog) GetActor() string {
	if y.Actor != nil {
		return *y.Actor
	}
	return ""
}
