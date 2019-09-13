package types

// SearchRoomsQueryVariables allowed parameters
type SearchRoomsQueryVariables struct {
	Amenities                  interface{} `json:"amenities"`
	EndedAt                    interface{} `json:"endedAt"`
	EndedTime                  interface{} `json:"endedTime"`
	EventType                  string      `json:"eventType"`
	RentType                   interface{} `json:"rentType"`
	Geocode                    interface{} `json:"geocode"`
	HasDirectReservationPlans  interface{} `json:"hasDirectReservationPlans"`
	HasLastMinuteDiscountPlans interface{} `json:"hasLastMinuteDiscountPlans"`
	HasTodayReservationPlans   interface{} `json:"hasTodayReservationPlans"`
	Keyword                    interface{} `json:"keyword"`
	Location                   interface{} `json:"location"`
	MaxCapacity                interface{} `json:"maxCapacity"`
	MaxPrice                   interface{} `json:"maxPrice"`
	MinCapacity                interface{} `json:"minCapacity"`
	MinPrice                   interface{} `json:"minPrice"`
	Page                       int         `json:"page"`
	PerPage                    int         `json:"perPage"`
	PriceType                  string      `json:"priceType"`
	SponsoredPromotionIds      interface{} `json:"sponsoredPromotionIds"`
	StartedAt                  interface{} `json:"startedAt"`
	StartedTime                interface{} `json:"startedTime"`
	State                      interface{} `json:"state"`
	WithRecommend              interface{} `json:"withRecommend"`
	Query                      string
}

// Space Room info
type RoomInfo struct {
	UID         string `json:"uid"`
	RentType    string `json:"rentType"`
	CanRentType string `json:"canRentType"`
	Query       string
}
