package golaget

import (
	"fmt"
)

type Store []StoreElement

type StoreElement struct {
	SiteID                     string        `json:"siteId"`
	Alias                      *string       `json:"alias"`
	IsActive                   bool          `json:"isActive"`
	IsBlocked                  bool          `json:"isBlocked"`
	IsOpen                     bool          `json:"isOpen"`
	IsBlockedByOrderLimit      bool          `json:"isBlockedByOrderLimit"`
	MaxOrdersPerDay            interface{}   `json:"maxOrdersPerDay"`
	OrdersToday                int64         `json:"ordersToday"`
	Address                    string        `json:"address"`
	PostalCode                 string        `json:"postalCode"`
	City                       string        `json:"city"`
	Phone                      *string       `json:"phone"`
	County                     string        `json:"county"`
	IsFullAssortmentOrderStore bool          `json:"isFullAssortmentOrderStore"`
	IsTastingStore             bool          `json:"isTastingStore"`
	Position                   *Position     `json:"position"`
	OpeningHours               []OpeningHour `json:"openingHours"`
	ParentSiteID               string        `json:"parentSiteId"`
	SearchArea                 *string       `json:"searchArea"`
}

// StoreGet
// Gets all stores
func (s SystemGoLaget) StoreGet() (response *Store, err error) {
	url := "https://api-extern.systembolaget.se/site/V2/Store"
	response = new(Store)
	s.JsonDecode(url, response)
	return
}

// StoreGetById
// Gets a specific store by id
func (s SystemGoLaget) StoreGetById(id string) (response *StoreElement, err error) {
	url := fmt.Sprintf("https://api-extern.systembolaget.se/site/V2/Store/%v", id)
	response = new(StoreElement)
	s.JsonDecode(url, response)
	return
}
