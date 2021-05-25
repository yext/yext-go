package yext

import "encoding/json"

const (
	ENTITYTYPE_HELPARTICLE EntityType = "helpArticle"
)

type HelpArticle struct {
	BaseEntity

	// Admin
	Keywords *[]string `json:"keywords,omitempty"` // LIST

	// Content
	Name             *string `json:"name,omitempty"`             // STRING
	Body             *string `json:"body,omitempty"`             // RICH_TEXT
	ShortDescription *string `json:"shortDescription,omitempty"` // STRING
	VoteCount        *int    `json:"voteCount,omitempty"`        // STRING
	VoteSum          *int    `json:"voteSum,omitempty"`          // STRING
	Promoted         **bool  `json:"promoted,omitempty"`         // BOOl

	ExternalArticlePostDate   *Date `json:"externalArticlePostDate,omitempty"`
	ExternalArticleUpdateDate *Date `json:"externalArticleUpdateDate,omitempty"`

	// Knowledge Assistant
	PrimaryConversationContact *string `json:"primaryConversationContact,omitempty"`
	NudgeEnabled               **bool  `json:"nudgeEnabled,omitempty"`

	// Internal Use
	Folder   *string `json:"folder,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

func (j *HelpArticle) String() string {
	b, _ := json.Marshal(j)
	return string(b)
}

func (h HelpArticle) GetId() string {
	if h.BaseEntity.Meta != nil && h.BaseEntity.Meta.Id != nil {
		return *h.BaseEntity.Meta.Id
	}
	return ""
}

func (h HelpArticle) GetName() string {
	if h.Name != nil {
		return *h.Name
	}
	return ""
}
