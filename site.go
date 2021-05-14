package golaget

import "fmt"

type Search struct {
	SearchModel               []SearchModel `json:"siteViewModel"`
	GooglePredictionViewModel []interface{} `json:"googlePredictionViewModel"`
}

type SearchModel struct {
	SiteID                string        `json:"siteId"`
	Alias                 *string       `json:"alias"`
	StreetAddress         string        `json:"streetAddress"`
	DisplayName           string        `json:"displayName"`
	City                  string        `json:"city"`
	County                string        `json:"county"`
	PostalCode            interface{}   `json:"postalCode"`
	IsTastingStore        bool          `json:"isTastingStore"`
	IsAgent               bool          `json:"isAgent"`
	IsOpen                bool          `json:"isOpen"`
	IsBlocked             bool          `json:"isBlocked"`
	IsBlockedByOrderLimit bool          `json:"isBlockedByOrderLimit"`
	OpeningHours          []OpeningHour `json:"openingHours"`
	Position              Position      `json:"position"`
}

// SiteSearch
// Searches for sites
func (s SystemGoLaget) SiteSearch(query string) (response *[]SearchModel, err error) {
	url := fmt.Sprintf("https://api-extern.systembolaget.se/site/V2/Search/Site/?q=%v", query)
	search := new(Search)
	s.JsonDecode(url, search)
	response = &search.SearchModel
	return
}

// TODO: Maybe include V1 API calls here?
// Or are those deprecated? They seem to be returning the same information as
// *Store
