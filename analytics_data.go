package yext

type AnalyticsData struct {
	ProfileViews                          *int     `json:"Profile Views"`
	Searches                              *int     `json:"Searches"`
	PowerlistingsLive                     *int     `json:"Powerlistings Live"`
	FeaturedMessageClicks                 *int     `json:"Featured Message Clicks"`
	YelpPageViews                         *int     `json:"Yelp Page Views"`
	BingSearches                          *int     `json:"Bing Searches"`
	FacebookLikes                         *int     `json:"Facebook Likes"`
	FacebookTalkingAbout                  *int     `json:"Facebook Talking About"`
	FacebookWereHere                      *int     `json:"Facebook Were Here"`
	FacebookCtaClicks                     *int     `json:"Facebook Cta Clicks"`
	FacebookImpressions                   *int     `json:"Facebook Impressions"`
	FacebookCheckins                      *int     `json:"Facebook Checkins"`
	FacebookPageViews                     *int     `json:"Facebook Page Views"`
	FacebookPostImpressions               *int     `json:"Facebook Post Impressions"`
	FoursquareDailyCheckins               *int     `json:"Foursquare Daily Checkins"`
	InstagramPosts                        *int     `json:"Instagram Posts"`
	GoogleQueryType                       *string  `json:"google_query_type"`
	GoogleSearchQueries                   *int     `json:"Google Search Queries"`
	GoogleSearchViews                     *int     `json:"Google Search Views"`
	GoogleMapViews                        *int     `json:"Google Maps Views"`
	GoogleCustomerActions                 *int     `json:"Google Customer Actions"`
	GooglePhoneCalls                      *int     `json:"Google Phone Calls"`
	YelpCustomerActions                   *int     `json:"Yelp Customer Actions"`
	AverageRating                         *float64 `json:"Average Rating"`
	NewReviews                            *int     `json:"Reviews"`
	StorepagesSessions                    *int     `json:"Storepages Sessions"`
	StorepagesPageviews                   *int     `json:"Page Views"`
	StorepagesDrivingdirections           *int     `json:"Driving Directions"`
	StorepagesPhonecalls                  *int     `json:"Taps to Call"`
	StorepagesCalltoactionclicks          *int     `json:"Call to Action Clicks"`
	StorepagesClickstowebsite             *int     `json:"Clicks to Website"`
	StorepagesEventEventtype              *int     `json:"Storepages Event Eventtype"`
	ProfileUpdates                        *int     `json:"Profile Updates"`
	PublisherSuggestions                  *int     `json:"Publisher Suggestions"`
	SocialActivities                      *int     `json:"Social Activities"`
	DuplicatesSuppressed                  *int     `json:"Duplicates Suppressed"`
	DuplicatesDetected                    *int     `json:"Duplicates Detected"`
	ListingsLive                          *int     `json:"Listings Live"`
	IstSearchRequests                     *int     `json:"Ist Search Requests"`
	IstAverageLocalPackPosition           *float64 `json:"Ist Average Local Pack Position"`
	IstAverageLocalPackNumberOfResults    *float64 `json:"Ist Average Local Pack Number Of Results"`
	IstLocalPackExisted                   *float64 `json:"Ist Local Pack Existed"`
	IstLocalPackPresence                  *float64 `json:"Ist Local Pack Presence"`
	IstKnowledgeCardExisted               *float64 `json:"Ist Knowledge Card Existed"`
	IstMatchesPerSearch                   *int     `json:"Ist Matches Per Search"`
	IstAverageFirstOrganicMatchPosition   *float64 `json:"Ist Average First Organic Match Position"`
	IstAverageFirstLocalPackMatchPosition *float64 `json:"Ist Average First Local Pack Match Position"`
	IstAverageFirstMatchPosition          *float64 `json:"Ist Average First Match Position"`
	IstOrganicShareOfSearch               *float64 `json:"Ist Organic Share Of Search"`
	IstLocalPackShareOfSearch             *float64 `json:"Ist Local Pack Share Of Search"`
	IstShareOfIntelligentSearch           *float64 `json:"Ist Share Of Intelligent Search"`
	LocationId                            *string  `json:"location_id"`
	Month                                 *string  `json:"month"`
	ResponseTime                          *int     `json:"Response Time (Hours)"`
	ResponseRate                          *int     `json:"Response Rate"`
	PartnerSite                           *string  `json:"site"`
	CumulativeRating                      *float64 `json:"Rolling Average Rating"`
	Competitor                            *string  `json:"competitor"`
}

func (y AnalyticsData) GetCompetitor() string {
	if y.Competitor != nil {
		return *y.Competitor
	}
	return ""
}

func (y AnalyticsData) GetProfileViews() int {
	if y.ProfileViews != nil {
		return *y.ProfileViews
	}
	return 0
}

func (y AnalyticsData) GetSearches() int {
	if y.Searches != nil {
		return *y.Searches
	}
	return 0
}

func (y AnalyticsData) GetPowerlistingsLive() int {
	if y.PowerlistingsLive != nil {
		return *y.PowerlistingsLive
	}
	return 0
}

func (y AnalyticsData) GetFeaturedMessageClicks() int {
	if y.FeaturedMessageClicks != nil {
		return *y.FeaturedMessageClicks
	}
	return 0
}

func (y AnalyticsData) GetYelpPageViews() int {
	if y.YelpPageViews != nil {
		return *y.YelpPageViews
	}
	return 0
}

func (y AnalyticsData) GetBingSearches() int {
	if y.BingSearches != nil {
		return *y.BingSearches
	}
	return 0
}

func (y AnalyticsData) GetFacebookLikes() int {
	if y.FacebookLikes != nil {
		return *y.FacebookLikes
	}
	return 0
}

func (y AnalyticsData) GetFacebookTalkingAbout() int {
	if y.FacebookTalkingAbout != nil {
		return *y.FacebookTalkingAbout
	}
	return 0
}

func (y AnalyticsData) GetFacebookWereHere() int {
	if y.FacebookWereHere != nil {
		return *y.FacebookWereHere
	}
	return 0
}

func (y AnalyticsData) GetFacebookCtaClicks() int {
	if y.FacebookCtaClicks != nil {
		return *y.FacebookCtaClicks
	}
	return 0
}

func (y AnalyticsData) GetFacebookImpressions() int {
	if y.FacebookImpressions != nil {
		return *y.FacebookImpressions
	}
	return 0
}

func (y AnalyticsData) GetFacebookCheckins() int {
	if y.FacebookCheckins != nil {
		return *y.FacebookCheckins
	}
	return 0
}

func (y AnalyticsData) GetFacebookPageViews() int {
	if y.FacebookPageViews != nil {
		return *y.FacebookPageViews
	}
	return 0
}

func (y AnalyticsData) GetFacebookPostImpressions() int {
	if y.FacebookPostImpressions != nil {
		return *y.FacebookPostImpressions
	}
	return 0
}

func (y AnalyticsData) GetFoursquareDailyCheckins() int {
	if y.FoursquareDailyCheckins != nil {
		return *y.FoursquareDailyCheckins
	}
	return 0
}

func (y AnalyticsData) GetInstagramPosts() int {
	if y.InstagramPosts != nil {
		return *y.InstagramPosts
	}
	return 0
}

func (y AnalyticsData) GetGoogleSearchQueries() int {
	if y.GoogleSearchQueries != nil {
		return *y.GoogleSearchQueries
	}
	return 0
}

func (y AnalyticsData) GetGoogleQueryType() string {
	if y.GoogleQueryType != nil {
		return *y.GoogleQueryType
	}
	return ""
}

func (y AnalyticsData) GetGoogleSearchViews() int {
	if y.GoogleSearchViews != nil {
		return *y.GoogleSearchViews
	}
	return 0
}

func (y AnalyticsData) GetGoogleMapViews() int {
	if y.GoogleMapViews != nil {
		return *y.GoogleMapViews
	}
	return 0
}

func (y AnalyticsData) GetGoogleCustomerActions() int {
	if y.GoogleCustomerActions != nil {
		return *y.GoogleCustomerActions
	}
	return 0
}

func (y AnalyticsData) GetGooglePhoneCalls() int {
	if y.GooglePhoneCalls != nil {
		return *y.GooglePhoneCalls
	}
	return 0
}

func (y AnalyticsData) GetYelpCustomerActions() int {
	if y.YelpCustomerActions != nil {
		return *y.YelpCustomerActions
	}
	return 0
}

func (y AnalyticsData) GetAverageRating() float64 {
	if y.AverageRating != nil {
		return *y.AverageRating
	}
	return -1
}

func (y AnalyticsData) GetCumulativeRating() float64 {
	if y.CumulativeRating != nil {
		return *y.CumulativeRating
	}
	return -1
}

func (y AnalyticsData) GetNewReviews() int {
	if y.NewReviews != nil {
		return *y.NewReviews
	}
	return 0
}

func (y AnalyticsData) GetStorepagesSessions() int {
	if y.StorepagesSessions != nil {
		return *y.StorepagesSessions
	}
	return 0
}

func (y AnalyticsData) GetStorepagesPageviews() int {
	if y.StorepagesPageviews != nil {
		return *y.StorepagesPageviews
	}
	return 0
}

func (y AnalyticsData) GetStorepagesDrivingdirections() int {
	if y.StorepagesDrivingdirections != nil {
		return *y.StorepagesDrivingdirections
	}
	return 0
}

func (y AnalyticsData) GetStorepagesPhonecalls() int {
	if y.StorepagesPhonecalls != nil {
		return *y.StorepagesPhonecalls
	}
	return 0
}

func (y AnalyticsData) GetStorepagesCalltoactionclicks() int {
	if y.StorepagesCalltoactionclicks != nil {
		return *y.StorepagesCalltoactionclicks
	}
	return 0
}

func (y AnalyticsData) GetStorepagesClickstowebsite() int {
	if y.StorepagesClickstowebsite != nil {
		return *y.StorepagesClickstowebsite
	}
	return 0
}

func (y AnalyticsData) GetStorepagesEventEventtype() int {
	if y.StorepagesEventEventtype != nil {
		return *y.StorepagesEventEventtype
	}
	return 0
}

func (y AnalyticsData) GetProfileUpdates() int {
	if y.ProfileUpdates != nil {
		return *y.ProfileUpdates
	}
	return 0
}

func (y AnalyticsData) GetPublisherSuggestions() int {
	if y.PublisherSuggestions != nil {
		return *y.PublisherSuggestions
	}
	return 0
}

func (y AnalyticsData) GetSocialActivities() int {
	if y.SocialActivities != nil {
		return *y.SocialActivities
	}
	return 0
}

func (y AnalyticsData) GetDuplicatesSuppressed() int {
	if y.DuplicatesSuppressed != nil {
		return *y.DuplicatesSuppressed
	}
	return 0
}

func (y AnalyticsData) GetDuplicatesDetected() int {
	if y.DuplicatesDetected != nil {
		return *y.DuplicatesDetected
	}
	return 0
}

func (y AnalyticsData) GetListingsLive() int {
	if y.ListingsLive != nil {
		return *y.ListingsLive
	}
	return 0
}

func (y AnalyticsData) GetIstSearchRequests() int {
	if y.IstSearchRequests != nil {
		return *y.IstSearchRequests
	}
	return 0
}

func (y AnalyticsData) GetIstAverageLocalPackPosition() float64 {
	if y.IstAverageLocalPackPosition != nil {
		return *y.IstAverageLocalPackPosition
	}
	return -1
}

func (y AnalyticsData) GetIstAverageLocalPackNumberOfResults() float64 {
	if y.IstAverageLocalPackNumberOfResults != nil {
		return *y.IstAverageLocalPackNumberOfResults
	}
	return -1
}

func (y AnalyticsData) GetIstLocalPackExisted() float64 {
	if y.IstLocalPackExisted != nil {
		return *y.IstLocalPackExisted
	}
	return -1
}

func (y AnalyticsData) GetIstLocalPackPresence() float64 {
	if y.IstLocalPackPresence != nil {
		return *y.IstLocalPackPresence
	}
	return -1
}

func (y AnalyticsData) GetIstKnowledgeCardExisted() float64 {
	if y.IstKnowledgeCardExisted != nil {
		return *y.IstKnowledgeCardExisted
	}
	return -1
}

func (y AnalyticsData) GetIstMatchesPerSearch() int {
	if y.IstMatchesPerSearch != nil {
		return *y.IstMatchesPerSearch
	}
	return 0
}

func (y AnalyticsData) GetIstAverageFirstOrganicMatchPosition() float64 {
	if y.IstAverageFirstOrganicMatchPosition != nil {
		return *y.IstAverageFirstOrganicMatchPosition
	}
	return -1
}

func (y AnalyticsData) GetIstAverageFirstLocalPackMatchPosition() float64 {
	if y.IstAverageFirstLocalPackMatchPosition != nil {
		return *y.IstAverageFirstLocalPackMatchPosition
	}
	return -1
}

func (y AnalyticsData) GetIstAverageFirstMatchPosition() float64 {
	if y.IstAverageFirstMatchPosition != nil {
		return *y.IstAverageFirstMatchPosition
	}
	return -1
}

func (y AnalyticsData) GetIstOrganicShareOfSearch() float64 {
	if y.IstOrganicShareOfSearch != nil {
		return *y.IstOrganicShareOfSearch
	}
	return -1
}

func (y AnalyticsData) GetIstLocalPackShareOfSearch() float64 {
	if y.IstLocalPackShareOfSearch != nil {
		return *y.IstLocalPackShareOfSearch
	}
	return -1
}

func (y AnalyticsData) GetIstShareOfIntelligentSearch() float64 {
	if y.IstShareOfIntelligentSearch != nil {
		return *y.IstShareOfIntelligentSearch
	}
	return -1
}

func (y AnalyticsData) GetLocationId() string {
	if y.LocationId != nil {
		return *y.LocationId
	}
	return ""
}

func (y AnalyticsData) GetMonth() string {
	if y.Month != nil {
		return *y.Month
	}
	return ""
}
