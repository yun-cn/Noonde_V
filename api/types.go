package api

import (
	"github.com/shopspring/decimal"
	"time"
)

// User ..
type User struct {
	ID        int64      `db:"id"                                         json:"id,omitempty"`         // User ID
	Email     string     `db:"email"                                      json:"email,omitempty"`      // Email
	Password  string     `db:"password"                                   json:"password,omitempty"`   // Password
	Nickname  string     `db:"nickname"                                   json:"nickname,omitempty"`   // Nickname
	Profile   string     `db:"profile"                                    json:"profile,omitempty"`    // Profile
	Avatar    string     `db:"avatar"                                     json:"avatar,omitempty"`     // S3 path for avatar
	AvatarSum string     `db:"avatar_sum"                                 json:"avatar_sum,omitempty"` // Check sum of avatar
	MaxTokens int        `db:"max_tokens"                                 json:"max_tokens,omitempty"` // Max token count
	Failed    int        `db:"failed"                                     json:"failed,omitempty"`     // Sign in failed count
	Created   *time.Time `db:"created"                                    json:"created,omitempty"`    // Created date
	Updated   *time.Time `db:"updated"                                    json:"updated,omitempty"`    // Updated date
}

// UserToken ..
type UserToken struct {
	ID      int64      `db:"id"                                           json:"id,omitempty"`      // User Token ID
	UserID  int64      `db:"user_id"                                      json:"user_id,omitempty"` // User ID
	Token   string     `db:"token"                                        json:"token,omitempty"`   // Token
	Expired *time.Time `db:"expired"                                      json:"synced,omitempty"`  // Expired date
	Created *time.Time `db:"created"                                      json:"created,omitempty"` // Created date
	Updated *time.Time `db:"updated"                                      json:"updated,omitempty"` // Updated date
}

// RelationShip..
type RelationShip struct {
	ID         int64      `db:"id"                                        json:"id,omitempty"`          // RelationShip
	FollowerID int32      `db:"follower_id"                               json:"follower_id,omitempty"` // FollowerID
	FollowedID int32      `db:"followed_id"                               json:"followed_id,omitempty"` // FollowedID
	Created    *time.Time `db:"created"                                   json:"created,omitempty"`     // Created date
	Updated    *time.Time `db:"updated"                                   json:"updated,omitempty"`     // Updated date
}

// Note ..
type Note struct {
	ID              int64      `db:"id"                                   json:"id,omitempty"`               // Note ID
	UserID          int64      `db:"user_id"                              json:"user_id,omitempty"`          // User ID
	UniqKey         string     `db:"uniq_key"                             json:"uniq_key,omitempty"`         // Uniq Key
	Name            string     `db:"name"                                 json:"name,omitempty"`             // Name
	Tags            string     `db:"tag"                                  json:"tag,omitempty"`              // Tags   例: 民泊　Space　...
	TweetText       string     `db:"tweet_text"                           json:"tweet_text,omitempty"`       // "ANA FIRST｜suzukyuin @suzukyuin｜note（ノート）"
	Body            string     `db:"body"                                 json:"body,omitempty"`             //  Body 内容
	LikeCount       int32      `db:"like_count"                           json:"like_count,omitempty"`       //  Like count
	CommentCount    int32      `db:"comment_count"                        json:"like_count,omitempty"`       //  Comment count
	CommentViewable int        `db:"comment_viewable"                     json:"comment_viewable,omitempty"` //  Comment viewable
	States          int        `db:"states"                               json:"states,omitempty"`           //  States
	DisplayDate     string     `db:"display_date"                         json:"states,omitempty"`           //  Display date
	CanRead         int        `db:"can_read"                             json:"can_read,omitempty"`         //  Can read
	Created         *time.Time `db:"created"                              json:"created,omitempty"`          // Created date
	Updated         *time.Time `db:"updated"                              json:"updated,omitempty"`          // Updated date
}

// NoteLike ..
type NoteLike struct {
	ID      int64      `db:"id"                                           json:"id,omitempty"`      //  NoteLike ID
	NoteID  int64      `db:"note_id"                                      json:"note_id,omitempty"` //  Note ID
	UserID  int64      `db:"user_id"                                      json:"user_id,omitempty"` //  User ID
	Created *time.Time `db:"created"                                      json:"created,omitempty"` // Created date
	Updated *time.Time `db:"updated"                                      json:"updated,omitempty"` // Updated date
}

// Comment ...
type Comment struct {
	ID        int64      `db:"id"                                         json:"id,omitempty"`         // Comment ID
	UserID    int64      `db:"user_id"                                    json:"user_id,omitempty"`    // User ID
	NoteID    int64      `db:"note_id"                                    json:"note_id,omitempty"`    // Note ID
	Comment   string     `db:"comment"                                    json:"comment,omitempty"`    // Comment
	LikeCount int32      `db:"like_count"                                 json:"like_count,omitempty"` // Like Count
	Created   *time.Time `db:"created"                                    json:"created,omitempty"`    // Created date
	Updated   *time.Time `db:"updated"                                    json:"updated,omitempty"`    // Updated date
}

// CommentLike ..
type CommentLike struct {
	ID        int64      `db:"id"                                         json:"id,omitempty"`         // CommentLike ID
	CommentID int64      `db:"comment_id"                                 json:"comment_id,omitempty"` // Comment ID
	UserID    int64      `db:"user_id"                                    json:"user_id,omitempty"`    // User ID
	States    int        `db:"states"                                     json:"states,omitempty"`     // States
	Created   *time.Time `db:"created"                                    json:"created,omitempty"`    // Created date
	Updated   *time.Time `db:"updated"                                    json:"updated,omitempty"`    // Updated date
}

// SpacemarketEventType ..
type SpacemarketEventType struct {
	ID        int        `db:"id"                                         json:"id,omitempty"`         // SpacemarketEventType ID
	State     int        `db:"state"                                      json:"state,omitempty"`      // State:  Ready: 1 Active: 2
	EventType string     `db:"event_type"                                 json:"event_type,omitempty"` // Event type
	StartPage int        `db:"start_page"                                 json:"start_page,omitempty"` // Start Page
	HourlyAt  *time.Time `db:"hourly_at"                                  json:"hourly_at,omitempty"`  // Hourly At
	DailyAt   *time.Time `db:"daily_at"                                   json:"daily_at,omitempty"`   // Daily At
	Created   *time.Time `db:"created"                                    json:"created,omitempty"`    // Created At
	Updated   *time.Time `db:"updated"                                    json:"updated,omitempty"`    // Update At
}

// InstabaseEventType
type InstabaseEventType struct {
	ID            int        `db:"id"                                     json:"id,omitempty"`              // Instabase's EventType ID
	State         int        `db:"state"                                  json:"state,omitempty"`           // Instabase's state
	EventTypeText string     `db:"event_type_text"                        json:"event_type_text,omitempty"` // Instabase's event type text
	EventTypeEn   string     `db:"event_type_en"                          json:"event_type_en,omitempty"`   // Instabase's event type english
	EventType     int        `db:"event_type"                             json:"event_type,omitempty"`      // Instabase's event type number
	StartPage     int        `db:"start_page"                             json:"start_page,omitempty"`      // Instabase's event start page
	HourlyAt      *time.Time `db:"hourly_at"                              json:"hourly_at,omitempty"`       // Instabase's houryly at
	Created       *time.Time `db:"created"                                json:"created,omitempty"`         // Created At
	Updated       *time.Time `db:"updated"                                json:"updated,omitempty"`         // Update At
}

// Platform
type Platform struct {
	ID         int        `db:"id"                                        json:"id,omitempty"`          // Platform id
	Url        string     `db:"url"                                       json:"url,omitempty"`         // Platform url
	ListingUrl string     `db:"listing_url"                               json:"listing_url,omitempty"` // Platform's listing url
	Created    *time.Time `db:"created"                                   json:"created,omitempty"`     // Created At
	Updated    *time.Time `db:"updated"                                   json:"updated,omitempty"`     // Update At
}

// Space ..
type Space struct {
	ID                   int64           `db:"id"                         json:"id,omitempty"`                 // Space Id
	Iop                  int64           `db:"iop"                        json:"iop,omitempty"`                // Id on platform
	Uip                  string          `db:"uip"                        json:"uip,omitempty"`                // Uid on platform
	Hiop                 int64           `db:"hiop"                       json:"hiop,omitempty"`               // Host Id on platform
	Name                 string          `db:"name"                       json:"name,omitempty"`               // Name on platform
	Capacity             int32           `db:"capacity"                   json:"capacity,omitempty"`           // Space's  guest capacity
	Description          string          `db:"description"                json:"description,omitempty"`        // Space's description
	EquipmentDescription string          `db:"equipment_description"      json:"equipment_description"`        // Space's equipment description
	Amenities            string          `db:"amenities"                  json:"amenities,omitempty"`          // Space's amenities
	EventTypes           string          `db:"event_types"                json:"event_types,omitempty"`        // Space's event types
	HourlyMinPrice       float64         `db:"hourly_min_price"           json:"hourly_min_price,omitempty"`   // Space's hourly min price
	HourlyMaxPrice       float64         `db:"hourly_max_price"           json:"hourly_max_price,omitempty"`   // Space's hourly max price
	DailyMinPrice        float64         `db:"daily_min_price"            json:"daily_min_price,omitempty"`    // Space's daily min price
	DailyMaxPrice        float64         `db:"daily_max_price"            json:"daily_max_price,omitempty"`    // Space's daily max price
	EmbedVideoUrl        string          `db:"embed_video_url"            json:"embed_video_url,omitempty"`    // Space's embed video url
	StateText            string          `db:"state_text"                 json:"state_text,omitempty"`         // Space's state text
	City                 string          `db:"city"                       json:"city,omitempty"`               // Space's city text
	Address              string          `db:"address"                    json:"address,omitempty"`            // Space's address
	Latitude             decimal.Decimal `db:"latitude"                   json:"latitude,omitempty"`           // Space's latitude
	Longitude            decimal.Decimal `db:"longitude"                  json:"longitude,omitempty"`          // Space's longitude
	Access               string          `db:"access"                     json:"access,omitempty"`             // Space's access
	Thumbnails           string          `db:"thumbnails"                 json:"thumbnails,omitempty"`         // Space's thumbnails
	ThirdReviewScore     decimal.Decimal `db:"third_review_score"         json:"third_review_score,omitempty"` // Space's third review score
	ThirdReviewCount     int32           `db:"third_review_count"         json:"third_review_count,omitempty"` // Space's third review count
	ThirdReplyRate       decimal.Decimal `db:"third_reply_rate"           json:"third_reply_rate,omitempty"`   // Space's third reply rate
	PlatformId           int             `db:"platform_id"                json:"platform_id,omitempty"`        // Space on platform id
	Hash                 string          `db:"hash"                       json:"hash,omitempty"`               // Space's hash
	HashAt               *time.Time      `db:"hash_at"                    json:"hash_at,omitempty"`            // Space's hash at
	SpaceURL             string          `db:"space_url"                  json:"space_url,omitempty"`          // Space's url
	States               int             `db:"states"                     json:"states, omitempty"`            // Space's states  { active: 1, inactive: 2 }
	Created              *time.Time      `db:"created"                    json:"created,omitempty"`            // Created At
	Updated              *time.Time      `db:"updated"                    json:"updated,omitempty"`            // Update At
}

// ..
const (
	TypeSpace = "space"
)

// ..
const (
	TableSpace = "spaces"
)
