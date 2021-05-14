package systemgolaget

import (
	"fmt"
)

type Agent []AgentElement

type AgentElement struct {
	AgentID                 string                  `json:"agentId"`
	Name                    string                  `json:"name"`
	Address                 string                  `json:"address"`
	PostalCode              string                  `json:"postalCode"`
	City                    string                  `json:"city"`
	Phone                   *string                 `json:"phone"`
	County                  *string                 `json:"county"`
	SearchArea              *string                 `json:"searchArea"`
	IsBlocked               *bool                   `json:"isBlocked"`
	IsActiveForPrepaidOrder bool                    `json:"isActiveForPrepaidOrder"`
	IsOpen                  bool                    `json:"isOpen"`
	IsBlockedByOrderLimit   bool                    `json:"isBlockedByOrderLimit"`
	MaxOrdersPerDay         *int64                  `json:"maxOrdersPerDay"`
	OrdersToday             int64                   `json:"ordersToday"`
	SiteID                  string                  `json:"siteId"`
	NextOrderStopTime       interface{}             `json:"nextOrderStopTime"`
	AgentDeliverySchedule   []AgentDeliverySchedule `json:"agentDeliverySchedule"`
	AgentPickupHours        []AgentPickupHour       `json:"agentPickupHours"`
	Position                *Position               `json:"position"`
}

type AgentDeliverySchedule struct {
	DeliveryDate  string `json:"deliveryDate"`
	OrderStopDate string `json:"orderStopDate"`
}

type AgentPickupHour struct {
	Date       string  `json:"date"`
	PickupFrom *string `json:"pickupFrom"`
	PickupTo   *string `json:"pickupTo"`
}

// AgentGet
// Gets all agents
func (s SystemGoLaget) AgentGet() (response *Agent, err error) {
	url := "https://api-extern.systembolaget.se/site/V2/Agent"
	response = new(Agent)
	s.JsonDecode(url, response)
	return
}

// AgentGetById
// Gets a specific agent by id
func (s SystemGoLaget) AgentGetById(id string) (response *AgentElement, err error) {
	url := fmt.Sprintf("https://api-extern.systembolaget.se/site/V2/Agent/%v", id)
	response = new(AgentElement)
	s.JsonDecode(url, response)
	return
}
