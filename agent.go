package golaget

import (
	"fmt"
)

// Agent is the initial reply from GetAgents
type Agent []AgentElement

// AgentElement contains all of the information gotten by the API
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

// AgentDeliverySchedule stores the dates the store does delivery
type AgentDeliverySchedule struct {
	DeliveryDate  string `json:"deliveryDate"`
	OrderStopDate string `json:"orderStopDate"`
}

// AgentPickupHour stores the dates that the store performs pickup
type AgentPickupHour struct {
	Date       string  `json:"date"`
	PickupFrom *string `json:"pickupFrom"`
	PickupTo   *string `json:"pickupTo"`
}

// AgentGet gets all agents
func (s SystemGoLaget) AgentGet() (response *Agent, err error) {
	url := "https://api-extern.systembolaget.se/site/V2/Agent"
	response = new(Agent)
	s.JsonDecode(url, response)
	return
}

// AgentGetById gets a specific AgentElement by id
func (s SystemGoLaget) AgentGetById(id string) (response *AgentElement, err error) {
	url := fmt.Sprintf("https://api-extern.systembolaget.se/site/V2/Agent/%v", id)
	response = new(AgentElement)
	s.JsonDecode(url, response)
	return
}
