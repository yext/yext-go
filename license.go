package yext

import "fmt"

type AssignedEntity struct {
	ExpirationDate string 	`json:"expirationDate,omitempty"`
	EntityId       string   `json:"entityId,omitempty"`
	LicensePackId  string   `json:"-"`
}

func (l *AssignedEntity) pathToAssignment() string {
	return fmt.Sprintf("%s/%s/assigned-entities/%s", packsPath, l.LicensePackId, l.EntityId)
}

type LicensePack struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Total     int    `json:"total,omitempty"`
	Assigned  int    `json:"assigned,omitempty"`
	Available int    `json:"available,omitempty"`
	Notes     string `json:"notes,omitempty"`
}

type LicenseAssignment struct {
	LicensePacks      *LicensePack     `json:"licensePack,omitempty"`
	AssignedEntities  []*AssignedEntity  `json:"assignedEntities,omitempty"`
}

type LicenseAssignmentListResponse struct {
	LicenseAssignments []*LicenseAssignment `json:"licenseAssignments,omitempty"`
	NextPageToken      string               `json:"nextPageToken,omitempty"`
}

type LicensePacksListResponse struct {
	LicensePacks   []*LicensePack  `json:"licensePacks,omitempty"`
	NextPageToken  string          `json:"nextPageToken,omitempty"`
}
