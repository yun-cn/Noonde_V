package types

// SearchRoomsQueryVariables allowed parameters
type SearchRoomsQueryVariables struct {
	AreaID         interface{} `json:"areaId"`
	BottomRightLat interface{} `json:"bottomRightLat"`
	BottomRightLon interface{} `json:"bottomRightLon"`
	CapacityIds    interface{} `json:"capacityIds"`
	CategoryIds    interface{} `json:"categoryIds"`
	ConditionIds   interface{} `json:"conditionIds"`
	EquipmentIds   interface{} `json:"equipmentIds"`
	FromDateDay    interface{} `json:"fromDateDay"`
	FromDateMonth  interface{} `json:"fromDateMonth"`
	FromDateYear   interface{} `json:"fromDateYear"`
	FromTime       interface{} `json:"fromTime"`
	OrderBy        interface{} `json:"orderBy"`
	Page           int         `json:"page"`
	PerPage        int         `json:"perPage"`
	PrefectureID   int         `json:"prefectureId"`
	StationID      interface{} `json:"stationId"`
	ToTime         interface{} `json:"toTime"`
	TopLeftLat     interface{} `json:"topLeftLat"`
	TopLeftLon     interface{} `json:"topLeftLon"`
	UsageIds       int         `json:"usageIds"`
	WardID         interface{} `json:"wardId"`
	Query          string
}

// Space details
type GetSpaceDetails struct {
	ID    int64  `json:"id"`
	Query string `json:"query"`
}
