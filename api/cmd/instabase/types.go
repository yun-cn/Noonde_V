package main

import (
	"github.com/shopspring/decimal"
	"time"
)

type briefSpace struct {
	ID               int64           `json:"id"`
	Name             string          `json:"name"`
	Capacity         int16           `json:"capacity"`
	ThirdReviewScore decimal.Decimal `json:"third_review_score"`
	ThirdReviewCount int32           `json:"third_review_count"`
	HourlyMinPrice   decimal.Decimal `json:"hourly_min_price"`
	HourlyMaxPrice   decimal.Decimal `json:"hourly_max_price"`
	SpaceURL         string          `json:"space_url"`
	Latitude         decimal.Decimal `json:"latitude"`
	Longitude        decimal.Decimal `json:"longitude"`
	Thumbnails       []byte          `json:"thumbnails"`
	PlatformId       int             `json:"platform_id"`
	Access           string          `json:"access"`
	Description      string          `json:"description"`
}

type spaceDetails struct {
	ID                   int64           `json:"id"`
	IOP                  int64           `json:"id"`
	UID                  string          `json:"uid"`
	Name                 string          `json:"name"`
	HIOP                 int64           `json:"hiop"`
	Capacity             int16           `json:"capacity"`
	Description          string          `json:"description"`
	EquipmentDescription string          `json:"equipment_description"`
	ReputationsCount     int             `json:"reputations_count"`
	Amenities            []byte          `json:"amenities"`
	EventTypes           []byte          `json:"eventTypes"`
	HourlyMinPrice       decimal.Decimal `json:"hourly_min_price"`
	HourlyMaxPrice       decimal.Decimal `json:"hourly_max_price"`
	DailyMinPrice        decimal.Decimal `json:"daily_min_price"`
	DailyMaxPrice        decimal.Decimal `json:"daily_max_price"`
	EmbedVideoUrl        string          `json:"embed_video_url"`
	StateText            string          `json:"state_text"`
	City                 string          `json:"city"`
	Address              string          `json:"address"`
	Latitude             decimal.Decimal `json:"latitude"`
	Longitude            decimal.Decimal `json:"longitude"`
	Access               string          `json:"access"`
	Hash                 string          `json:"hash"`
	HashAt               *time.Time      `json:"hash_at"`
	Thumbnails           []byte          `json:"thumbnails"`
	ThirdReplyRate       decimal.Decimal `json:"third_reply_rate"`
	ThirdReviewScore     decimal.Decimal `json:"third_review_score"`
	ThirdReviewCount     int32           `json:"third_review_count"`
	SpaceURL             string          `json:"space_url"`
	States               int             `json:"states"`
	PlatformId           int             `json:"platform_id"`
}
