package main

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"github.com/yun313350095/Noonde/api"
	"strconv"
	"strings"
	"time"
)

func (s *service) spaceRoomListHourly(client api.SpaceMarketClient, page, perPage int, location, eventType string) ([]*briefSpace, error) {
	var err error
	var body string
	var data gjson.Result

	// Get space
	body, err = client.SearchRoomListHourly(page, perPage, location, eventType)
	if err != nil {
		s.Log.Error(err)
		panic(err)
	}

	// Parse Space
	data = gjson.Parse(body)
	//s.Log.Info(data)
	records := data.Get("data.searchRooms.results")

	spaces := []*briefSpace{}

	// 翻页完成后
	if len(records.Array()) == 0 {
		return spaces, nil
	}

	for _, rec := range records.Array() {

		thumbnails := rec.Get("thumbnails.#.url")
		thumbnailList := make([]string, len(thumbnails.Array()))

		for i, thumbnail := range thumbnails.Array() {
			thumbnailList[i] = thumbnail.String()
		}

		thumbnailsArray, _ := json.Marshal(thumbnailList)

		// Hourly min price
		hourlyMinPriceStr := strings.TrimPrefix(rec.Get("prices.0.minText").String(), "¥")
		hourlyMinPriceStr = strings.Replace(hourlyMinPriceStr, ",", "", -1)
		hourlyMinPrice, _ := decimal.NewFromString(hourlyMinPriceStr)

		// Hourly max price reputation summary
		hourlyMaxPriceStr := strings.TrimPrefix(rec.Get("prices.0.maxText").String(), "¥")
		hourlyMaxPriceStr = strings.Replace(hourlyMaxPriceStr, ",", "", -1)
		hourlyMaxPrice, _ := decimal.NewFromString(hourlyMaxPriceStr)

		// Third review score
		thirdReviewScore, _ := decimal.NewFromString(rec.Get("rentalReputationScore").String())

		// Latitude and longitude
		latitude, _ := decimal.NewFromString(rec.Get("latitude").String())
		longitude, _ := decimal.NewFromString(rec.Get("longitude").String())

		// Get capacity and setup 1 people
		capacity := int16(1)
		if rec.Get("capacity").Int() != 0 {
			capacity = int16(rec.Get("capacity").Int())
		}

		spaces = append(spaces, &briefSpace{
			PlatformId:            1,
			Name:                  rec.Get("name").String(),
			ID:                    rec.Get("id").Int(),
			UID:                   rec.Get("uid").String(),
			Capacity:              capacity,
			RentalReputationScore: thirdReviewScore,
			TotalReputationCount:  int32(rec.Get("totalReputationCount").Int()),
			StateText:             rec.Get("stateText").String(),
			City:                  rec.Get("city").String(),
			Latitude:              latitude,
			Longitude:             longitude,
			Access:                rec.Get("access").String(),
			Thumbnails:            thumbnailsArray,
			HourlyMinPrice:        hourlyMinPrice,
			HourlyMaxPrice:        hourlyMaxPrice,
		})
	}

	return spaces, nil
}

func (s *service) saveSpaces(ctx context.Context, tx *sqlx.Tx, spaces []*briefSpace, oneType api.SpacemarketEventType) error {

	if len(spaces) >= 1 {
		for _, space := range spaces {
			// insert space if not exists.
			res, err := tx.ExecContext(ctx,
				`insert into spaces (
						iop,
						uip,
						name, 
						capacity,
						hourly_min_price,
						hourly_max_price,
						state_text,
						city,
						latitude,
						longitude,
						access,
						thumbnails,
						third_review_score,
						third_review_count,
						platform_id,
						created,
						updated,
						description,
						equipment_description,
						amenities,
						event_types) values(
						?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?
						) on duplicate key update 
						id = last_insert_id(id),
						updated = values(updated)`,
				space.ID,
				space.UID,
				space.Name,
				space.Capacity,
				space.HourlyMinPrice,
				space.HourlyMaxPrice,
				space.StateText,
				space.City,
				space.Latitude,
				space.Longitude,
				space.Access,
				space.Thumbnails,
				space.RentalReputationScore,
				space.TotalReputationCount,
				space.PlatformId,
				time.Now(),
				time.Now(),
				"",
				"",
				"",
				"")

			if err != nil {
				s.Log.Error(err)
				return err
			}

			// Update eventType
			spaceID, err := res.LastInsertId()
			s.Log.Info("Space ID:    " + strconv.FormatInt(spaceID, 10))
			s.Log.Info("Space Name:  " + space.Name)
			if err != nil {
				s.Log.Error(err)
				return err
			}

			//Update elastic ..
			var targets []*api.ElasticTarget
			targets = s.Elastic.AddTargets(targets, api.TypeSpace, []int64{spaceID})
			err = s.Elastic.BulkUpdate(ctx, tx, targets)
			if err != nil {
				s.Log.Error(err)
				return err
			}
		}

		_, err := tx.ExecContext(ctx, `Update spacemarket_event_types set start_page = ?,  hourly_at = now(), updated=now() where id = ?`,
			oneType.StartPage+1,
			oneType.ID)

		if err != nil {
			s.Log.Info(err)
			return err
		}
	} else {
		_, err := tx.ExecContext(ctx, `update spacemarket_event_types set start_page = ?, hourly_at = now(), updated = now() where id = ?`,
			1, oneType.ID)

		if err != nil {
			s.Log.Info(err)
			return err
		}
	}
	return nil
}
