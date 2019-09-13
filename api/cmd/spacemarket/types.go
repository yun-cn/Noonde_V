package main

import (
	"github.com/shopspring/decimal"
	"time"
)

type briefSpace struct {
	ID                         int64           `json:"id"`
	UID                        string          `json:"uid"`
	Name                       string          `json:"name"`
	HasDirectReservationPlans  bool            `json:"has_direct_reservation_plans"`
	HasLastMinuteDiscountPlans bool            `json:"has_last_minute_discount_plans"`
	HasTodayReservationPlans   bool            `json:"has_today_reservation_plans"`
	Capacity                   int16           `json:"capacity"`
	RentalReputationScore      decimal.Decimal `json:"rental_reputation_score"`
	RentalReputationCount      int32           `json:"rental_reputation_count"`
	TotalReputationScore       decimal.Decimal `json:"total_reputation_score"`
	TotalReputationCount       int32           `json:"total_reputation_count"`
	HourlyMinPrice             decimal.Decimal `json:"hourly_min_price"`
	HourlyMaxPrice             decimal.Decimal `json:"hourly_max_price"`
	DailyMinPrice              decimal.Decimal `json:"daily_min_price"`
	DailyMaxPrice              decimal.Decimal `json:"daily_max_price"`
	Thumbnails                 []byte          `json:"thumbnails"`
	IsInquiryOnly              bool            `json:"is_inquiry_only"`
	OwnerRank                  int             `json:"owner_rank"`
	StateText                  string          `json:"state_text"`
	City                       string          `json:"city"`
	Latitude                   decimal.Decimal `json:"latitude"`
	Longitude                  decimal.Decimal `json:"longitude"`
	SpaceUsername              string          `json:"space_username"`
	IsFavorite                 bool            `json:"is_favorite"`
	Access                     string          `json:"access"`
	AvailablePlanCount         int             `json:"available_plan_count"`
	Plans                      []interface{}   `json:"plans"`
	IsSponsoredPromotionRoom   bool            `json:"is_sponsored_promotion_room"`
	SponsoredPromotions        []interface{}   `json:"sponsored_promotions"`
	IsCancelFree               bool            `json:"is_cancel_free"`
	Typename                   string          `json:"__typename"`
	PlatformId                 int             `json:"platform_id"`
}

type spaceDetails struct {
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
	States               int             `json:"states"`
	PlatformId           int             `json:"platform_id"`
	SpaceURL             string          `json:"space_url"`
}
