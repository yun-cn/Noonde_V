package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"github.com/yun313350095/Noonde/api"
	"strconv"
	"strings"
	"time"
)

func (s *service) spaceRoomDetails(client api.SpaceMarketClient, uid string, iop int64, rentType string) (*spaceDetails, error) {
	var err error
	var body string
	var data gjson.Result

	// Get space info
	body, err = client.GetRoomDetails(uid, rentType)
	if err != nil {
		s.Log.Error(err)
		panic(err)
	}

	// Parse space
	data = gjson.Parse(body)
	record := data.Get("data")

	// Response body to hash
	hashByte := sha256.Sum256([]byte(body))
	hash := hex.EncodeToString(hashByte[:])

	// Edited  hash at
	hashAt := time.Now()
	s.Log.Info(record.Get(`room`).Exists())

	if !record.Get(`room`).Exists() {
		space := &spaceDetails{
			PlatformId: 1,
			States:     2,
			IOP:        iop,
			Hash:       hash,
			HashAt:     &hashAt,
		}

		return space, nil
	}

	// Thumbnails array
	thumbnails := record.Get("room.thumbnails.results.#.url")
	thumbnailList := make([]string, len(thumbnails.Array()))

	for i, thumbnail := range thumbnails.Array() {
		thumbnailList[i] = thumbnail.String()
	}

	thumbnailsArray, _ := json.Marshal(thumbnailList)

	// Amenities array
	amenities := record.Get("room.amenities.results.#.name")
	amenitiesList := make([]string, len(amenities.Array()))

	for i, amenitie := range amenities.Array() {
		amenitiesList[i] = amenitie.String()
	}

	amenitiesArray, _ := json.Marshal(amenitiesList)

	// Event types array
	eventTypes := record.Get("room.eventTypes.results.#.name")
	eventTypeList := make([]string, len(eventTypes.Array()))

	for i, eventType := range eventTypes.Array() {
		eventTypeList[i] = eventType.String()
	}
	eventTypesArray, _ := json.Marshal(eventTypeList)

	// Get hourly min price
	resHourlyMinPrice := record.Get(`room.prices.#[maxUnitText=="時間"].minText`)

	hourlyMinPrice := decimal.NewFromFloat(0.0)
	if strings.Contains(resHourlyMinPrice.String(), "￥") {
		hourlyMinPriceStr := strings.TrimPrefix(resHourlyMinPrice.String(), "￥")
		hourlyMinPriceStr = strings.Replace(hourlyMinPriceStr, ",", "", -1)
		hourlyMinPrice, _ = decimal.NewFromString(hourlyMinPriceStr)
	}

	// Get hourly max price
	resHourlyMaxPrice := record.Get(`room.prices.#[maxUnitText=="時間"].maxText`)
	hourlyMaxPrice := decimal.NewFromFloat(0.0)
	if strings.Contains(resHourlyMaxPrice.String(), "￥") {
		hourlyMaxPriceStr := strings.TrimPrefix(resHourlyMaxPrice.String(), "￥")
		hourlyMaxPriceStr = strings.Replace(hourlyMaxPriceStr, ",", "", -1)
		hourlyMaxPrice, _ = decimal.NewFromString(hourlyMaxPriceStr)
	}

	// Get daily min price
	resDailyMinPrice := record.Get(`room.prices.#[maxUnitText==日"].minText`)

	dailyMinPrice := decimal.NewFromFloat(0.0)
	if strings.Contains(resDailyMinPrice.String(), "￥") {
		dailyMinPriceStr := strings.TrimPrefix(resDailyMinPrice.String(), "￥")
		dailyMinPriceStr = strings.Replace(dailyMinPriceStr, ",", "", -1)
		dailyMinPrice, _ = decimal.NewFromString(dailyMinPriceStr)
	}

	// Get daily max price
	resDailyMaxPrice := record.Get(`room.prices.#[maxUnitText==日"].maxText`)
	dailyMaxPrice := decimal.NewFromFloat(0.0)
	if strings.Contains(resDailyMaxPrice.String(), "￥") {
		dailyMaxPriceStr := strings.TrimPrefix(resDailyMaxPrice.String(), "￥")
		dailyMaxPriceStr = strings.Replace(dailyMaxPriceStr, ",", "", -1)
		dailyMaxPrice, _ = decimal.NewFromString(dailyMaxPriceStr)
	}

	// Third review score
	thirdReviewScore, _ := decimal.NewFromString(record.Get(`room.reputationSummary.score`).String())

	// Third review count
	thirdReviewCount := int32(record.Get(`room.reputationSummary.count`).Int())

	// Third reply rate
	thirdReplyRate, _ := decimal.NewFromString(record.Get(`room.owner.replyRate`).String())

	// Latitude and longitude
	latitude, _ := decimal.NewFromString(record.Get(`room.space.latitude`).String())
	longitude, _ := decimal.NewFromString(record.Get(`room.space.longitude`).String())

	// Get capacity or setup 1 people
	capacity := int16(1)
	if record.Get(`room.capacity`).Int() != 0 {
		capacity = int16(record.Get(`room.capacity`).Int())
	}

	// Get space url
	username := record.Get(`room.space.username`).String()
	spaceUrl := "https://www.spacemarket.com/spaces/" + username + "/rooms/" + uid

	space := &spaceDetails{
		PlatformId:           1,
		States:               1,
		Name:                 record.Get(`room.name`).String(),
		IOP:                  record.Get(`room.id`).Int(),
		UID:                  record.Get(`room.uid`).String(),
		Capacity:             capacity,
		ThirdReplyRate:       thirdReplyRate,
		ThirdReviewCount:     thirdReviewCount,
		ThirdReviewScore:     thirdReviewScore,
		StateText:            record.Get(`room.space.stateText`).String(),
		City:                 record.Get(`room.space.city`).String(),
		Latitude:             latitude,
		Longitude:            longitude,
		Access:               record.Get(`room.space.access`).String(),
		Thumbnails:           thumbnailsArray,
		Amenities:            amenitiesArray,
		Address:              record.Get(`room.space.address1`).String() + record.Get(`room.space.address2`).String(),
		HourlyMinPrice:       hourlyMinPrice,
		HourlyMaxPrice:       hourlyMaxPrice,
		DailyMinPrice:        dailyMinPrice,
		DailyMaxPrice:        dailyMaxPrice,
		Hash:                 hash,
		HashAt:               &hashAt,
		Description:          record.Get(`room.description`).String(),
		HIOP:                 record.Get(`room.owner.id`).Int(),
		EmbedVideoUrl:        record.Get(`room.embedVideoUrl`).String(),
		EquipmentDescription: record.Get(`room.equipmentDescription`).String(),
		EventTypes:           eventTypesArray,
		SpaceURL:             spaceUrl,
	}

	return space, nil
}

func (s *service) saveRoomDetails(ctx context.Context, tx *sqlx.Tx, space *spaceDetails) error {

	// Space hash dose not exits or not same ?
	var hash string

	err := tx.GetContext(ctx, &hash, `select hash from spaces where iop = ? and platform_id = 1`, space.IOP)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	if (hash != space.Hash && space.States == 1) || (len(hash) == 0 && space.States == 1) {
		res, err := tx.ExecContext(ctx,
			`update spaces set
					uip = ?,
					hiop = ?,
					name = ?,
					capacity = ?,
					hourly_min_price = ?,
					hourly_max_price = ?,
					daily_min_price = ?,
					daily_max_price = ?,
					third_reply_rate = ?,
					third_review_count = ?,
					third_review_score = ?,
					state_text = ?,
					city = ?,
					latitude = ?,
					longitude = ?,
					access = ?,
					thumbnails = ?,
					amenities = ?,
					address = ?,
					hash = ?,
					hash_at = ?,
					description = ?,
					embed_video_url = ?,
					equipment_description = ?,
					event_types = ?,
					space_url = ?,
					updated = now() where platform_id = ? and iop = ?`,
			space.UID,
			space.HIOP,
			space.Name,
			space.Capacity,
			space.HourlyMinPrice,
			space.HourlyMaxPrice,
			space.DailyMinPrice,
			space.DailyMaxPrice,
			space.ThirdReplyRate,
			space.ThirdReviewCount,
			space.ThirdReviewScore,
			space.StateText,
			space.City,
			space.Latitude,
			space.Longitude,
			space.Access,
			space.Thumbnails,
			space.Amenities,
			space.Address,
			space.Hash,
			time.Now(),
			space.Description,
			space.EmbedVideoUrl,
			space.EquipmentDescription,
			space.EventTypes,
			space.SpaceURL,
			space.PlatformId,
			space.IOP)

		if err != nil {
			s.Log.Error(err)
			return err
		}

		updatedRow, err := res.RowsAffected()
		s.Log.Info(updatedRow)

		if err != nil {
			s.Log.Error(err)
			return err
		}
		var spaceID int64
		err = tx.GetContext(ctx, &spaceID, `select id from spaces where iop = ? and platform_id = 1`,
			space.IOP)

		if err != nil {
			s.Log.Error(err)
			return err
		}
		s.Log.Info("Space ID:    " + strconv.FormatInt(spaceID, 10))

		// Update  elastic_search
		var targets []*api.ElasticTarget
		targets = s.Elastic.AddTargets(targets, api.TypeSpace, []int64{spaceID})
		err = s.Elastic.BulkUpdate(ctx, tx, targets)
		if err != nil {
			s.Log.Error(err)
			return err
		}
	} else {
		res, err := tx.ExecContext(ctx,
			`update spaces set states = ?, hash = ?, hash_at = ? where iop = ? and platform_id = 1`,
			space.States,
			space.Hash,
			space.HashAt,
			space.IOP)

		if err != nil {
			s.Log.Error(err)
			return err
		}

		updatedRow, err := res.RowsAffected()
		if err != nil {
			s.Log.Error(err)
			return err
		}
		s.Log.Info(updatedRow)

		var spaceID int64
		err = tx.GetContext(ctx, &spaceID, `select id from spaces where iop = ? and platform_id = 1`,
			space.IOP)
		if err != nil {
			s.Log.Error(err)
			return err
		}
		s.Log.Info("Space ID:    " + strconv.FormatInt(spaceID, 10))

		// Update elastic_search
		var targets []*api.ElasticTarget
		targets = s.Elastic.AddTargets(targets, api.TypeSpace, []int64{spaceID})
		err = s.Elastic.BulkUpdate(ctx, tx, targets)
		if err != nil {
			s.Log.Error(err)
			return err
		}
	}

	return nil
}
