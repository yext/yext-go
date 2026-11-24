package yext

import (
	"encoding/json"
)

const ENTITYTYPE_HEALTHCAREPROVIDER EntityType = "healthcareProvider"

// HealthcareProviderEntity is the representation of a Healthcare Provider: the person-level entity in the
// provider / facility data model. Data that is true of a provider regardless of where they practice lives here;
// data specific to a provider at a given facility lives on ProviderFacilityEntity, which links back through its
// healthcareProvider field.
type HealthcareProviderEntity struct {
	BaseEntity

	// Linked Provider Facilities (entity IDs)
	ProviderFacilityRefs *[]string `json:"providerFacilityRefs,omitempty"`

	// Name / Identity
	Name       *string           `json:"name,omitempty"`
	FirstName  *string           `json:"firstName,omitempty"`
	MiddleName *string           `json:"middleName,omitempty"`
	LastName   *string           `json:"lastName,omitempty"`
	Gender     **string          `json:"gender,omitempty"`
	NPI        *string           `json:"npi,omitempty"`
	Degrees    *UnorderedStrings `json:"degrees,omitempty"`
	Slug       *string           `json:"slug,omitempty"`
	Headshot   **Image           `json:"headshot,omitempty"`

	// Profile
	Description           *string           `json:"description,omitempty"`
	Languages             *[]string         `json:"languages,omitempty"`
	Services              *[]string         `json:"services,omitempty"`
	Brands                *[]string         `json:"brands,omitempty"`
	Associations          *[]string         `json:"associations,omitempty"`
	Certifications        *[]string         `json:"certifications,omitempty"`
	ConditionsTreated     *[]string         `json:"conditionsTreated,omitempty"`
	AdmittingHospitals    *[]string         `json:"admittingHospitals,omitempty"`
	EducationList         *[]Education      `json:"educationList,omitempty"`
	AcceptingNewPatients  *bool             `json:"acceptingNewPatients,omitempty"`
	TelehealthUrl         *string           `json:"telehealthUrl,omitempty"`
	Emails                *[]string         `json:"emails,omitempty"`
	COVID19InformationUrl *string           `json:"covid19InformationUrl,omitempty"`
	FeaturedMessage       **FeaturedMessage `json:"featuredMessage,omitempty"`

	// Media
	Logo         **Photo  `json:"logo,omitempty"`
	PhotoGallery *[]Photo `json:"photoGallery,omitempty"`
	Videos       *[]Video `json:"videos,omitempty"`

	// Social Media / Google / Facebook
	FacebookCoverPhoto   **Image `json:"facebookCoverPhoto,omitempty"`
	FacebookName         *string `json:"facebookName,omitempty"`
	FacebookPageUrl      *string `json:"facebookPageUrl,omitempty"`
	FacebookParentPageId *string `json:"facebookParentPageId,omitempty"`
	FacebookProfilePhoto **Image `json:"facebookProfilePhoto,omitempty"`
	GoogleCoverPhoto     **Image `json:"googleCoverPhoto,omitempty"`
	GoogleProfilePhoto   **Image `json:"googleProfilePhoto,omitempty"`
	GoogleShortName      *string `json:"googleShortName,omitempty"`
	InstagramHandle      *string `json:"instagramHandle,omitempty"`
	TwitterHandle        *string `json:"twitterHandle,omitempty"`
	LandingPageUrl       *string `json:"landingPageUrl,omitempty"`
	LinkedInUrl          *string `json:"linkedInUrl,omitempty"`
	XingUrl              *string `json:"xingUrl,omitempty"`
	// BrandReference holds the entity ID of the linked Brand entity.
	BrandReference    *string `json:"brandReference,omitempty"`
	YouTubeChannelUrl *string `json:"youTubeChannelUrl,omitempty"`

	// Apple
	AppleBusinessDescription *string `json:"appleBusinessDescription,omitempty"`
	AppleBusinessId          *string `json:"appleBusinessId,omitempty"`
	AppleCompanyId           *string `json:"appleCompanyId,omitempty"`
	AppleCoverPhoto          **Image `json:"appleCoverPhoto,omitempty"`
	AppleDisplayName         *string `json:"appleDisplayName,omitempty"`
}

func (p *HealthcareProviderEntity) UnmarshalJSON(data []byte) error {
	type Alias HealthcareProviderEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(p, data)
}

func (y HealthcareProviderEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y HealthcareProviderEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y HealthcareProviderEntity) GetName() string {
	if y.Name != nil {
		return *y.Name
	}
	return ""
}

func (y HealthcareProviderEntity) GetNPI() string {
	if y.NPI != nil {
		return *y.NPI
	}
	return ""
}

func (y HealthcareProviderEntity) GetProviderFacilityRefs() (v []string) {
	if y.ProviderFacilityRefs != nil {
		v = *y.ProviderFacilityRefs
	}
	return v
}

func (y HealthcareProviderEntity) GetLanguages() (v []string) {
	if y.Languages != nil {
		v = *y.Languages
	}
	return v
}

func (y HealthcareProviderEntity) GetAcceptingNewPatients() bool {
	return y.AcceptingNewPatients != nil && *y.AcceptingNewPatients
}
