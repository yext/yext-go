package yext

import (
	"encoding/json"
)

const ENTITYTYPE_JOB EntityType = "job"

type JobEntity struct {
	BaseEntity

	//Job Info
	Name           *string      `json:"name,omitempty"`
	Description    *string      `json:"description,omitempty"`
	Timezone       *string      `json:"timezone,omitempty"`
	EmploymentType *string      `json:"employmentType,omitempty"`
	DatePosted     *string      `json:"datePosted,omitempty"`
	ValidThrough   *string      `json:"validThrough,omitempty"`
	Keywords       *[]string    `json:"keywords,omitempty"`
	Location       *JobLocation `json:"location,omitempty"`

	// Urls
	ApplicationURL *string `json:"applicationUrl,omitempty"`
	LandingPageURL *string `json:"landingPageUrl,omitempty"`
}

type JobLocation struct {
	ExistingLocation *string `json:"existingLocation,omitempty"`
	ExternalLocation *string `json:"externalLocation,omitempty"`
}

func (j JobEntity) GetId() string {
	if j.BaseEntity.Meta != nil && j.BaseEntity.Meta.Id != nil {
		return *j.BaseEntity.Meta.Id
	}
	return ""
}

func (j JobEntity) GetName() string {
	if j.Name != nil {
		return *j.Name
	}
	return ""
}

func (j JobEntity) GetDescription() string {
	if j.Description != nil {
		return *j.Description
	}
	return ""
}

func (j JobEntity) GetTimezone() string {
	if j.Timezone != nil {
		return *j.Timezone
	}
	return ""
}

func (j JobEntity) GetEmploymentType() string {
	if j.EmploymentType != nil {
		return *j.EmploymentType
	}
	return ""
}

func (j JobEntity) GetDatePosted() string {
	if j.DatePosted != nil {
		return *j.DatePosted
	}
	return ""
}

func (j JobEntity) GetExpiryDate() string {
	if j.ValidThrough != nil {
		return *j.ValidThrough
	}
	return ""
}

func (j JobEntity) GetApplicationURL() string {
	if j.ApplicationURL != nil {
		return *j.ApplicationURL
	}
	return ""
}

func (j JobEntity) GetLandingPageUrl() string {
	if j.LandingPageURL != nil {
		return *j.LandingPageURL
	}
	return ""
}

func (j JobEntity) GetExistingLocation() string {
	if j.Location != nil && j.Location.ExistingLocation != nil {
		return *j.Location.ExistingLocation
	}
	return ""
}

func (j JobEntity) GetExternalLocation() string {
	if j.Location != nil && j.Location.ExternalLocation != nil {
		return *j.Location.ExternalLocation
	}
	return ""
}

func (j *JobEntity) String() string {
	b, _ := json.Marshal(j)
	return string(b)
}
