package yext

const (
	activityLogPath = "analytics/activity"

	// Use the following constants for Activity Filter
	ACTIVITYTYPES_LOCATION_UPDATED              = "LOCATION_UPDATED"
	ACTIVITYTYPES_PUBLISHER_SUGGESTION_CREATED  = "PUBLISHER_SUGGESTION_CREATED"
	ACTIVITYTYPES_PUBLISHER_SUGGESTION_APPROVED = "PUBLISHER_SUGGESTION_APPROVED"
	ACTIVITYTYPES_PUBLISHER_SUGGESTION_REJECTED = "PUBLISHER_SUGGESTION_REJECTED"
	ACTIVITYTYPES_REVIEW_CREATED                = "REVIEW_CREATED"
	ACTIVITYTYPES_SOCIAL_POST_CREATED           = "SOCIAL_POST_CREATED"
	ACTIVITYTYPES_LISTING_LIVE                  = "LISTING_LIVE"
	ACTIVITYTYPES_DUPLICATE_SUPPRESSED          = "DUPLICATE_SUPPRESSED"

	ACTORS_YEXT_SYSTEM       = "YEXT_SYSTEM"
	ACTORS_SCHEDULED_CONTENT = "SCHEDULED_CONTENT"
	ACTORS_API               = "API"
	ACTORS_PUBLISHER         = "PUBLISHER"
)

type ActivityLogService struct {
	client *Client
}

type ActivityLogResponse struct {
	Count *int           `json:"count"`
	Data  []*ActivityLog `json:"activities"`
}

type ActivityFilter struct {
	StartDate     *string   `json:"startDate"`
	EndDate       *string   `json:"endDate"`
	LocationIds   *[]string `json:"locationIds"`
	ActivityTypes *[]string `json:"activityTypes"` // See const above
	Actors        *[]string `json:"actors"`        // See const above
}

type ActivityLogRequest struct {
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
	Filters *ActivityFilter `json:"filters"`
}

// ActivityLogRequest can be nil, no filter
func (a *ActivityLogService) Create(req *ActivityLogRequest) (*ActivityLogResponse, *Response, error) {
	arr := &ActivityLogResponse{}
	r, err := a.client.DoRequestJSON("POST", activityLogPath, req, arr)
	return arr, r, err
}
