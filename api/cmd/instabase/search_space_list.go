package main

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"github.com/yun313350095/Noonde/api"
	"strconv"
	"time"
)

func (s *service) searchSpaceList(client api.InstabaseClient, page int, eventType int) ([]*briefSpace, error) {
	var err error
	var body string
	var data gjson.Result

	// Get space
	body, err = client.SearchSpaceList(page, eventType)
	if err != nil {
		s.Log.Info(err)
		panic(err)
	}

	// Parse space
	data = gjson.Parse(body)

	records := data.Get("data.spaces")

	spaces := []*briefSpace{}

	if len(records.Array()) == 0 {
		return spaces, nil
	}

	for _, rec := range records.Array() {

		thumbnails := rec.Get("images.#.largeFilePath")
		thumbnailList := make([]string, len(thumbnails.Array()))

		for i, thumbnail := range thumbnails.Array() {
			thumbnailList[i] = thumbnail.String()
		}

		thumbnailsArray, _ := json.Marshal(thumbnailList)

		// Hourly min price
		hourlyMinPriceStr := rec.Get("summaryMinPrice").String()
		hourlyMinPrice, _ := decimal.NewFromString(hourlyMinPriceStr)
		s.Log.Info(hourlyMinPrice)

		// Hourly max price
		hourlyMaxPriceStr := rec.Get("summaryMaxPrice").String()
		hourlyMaxPrice, _ := decimal.NewFromString(hourlyMaxPriceStr)
		s.Log.Info(hourlyMaxPrice)

		// Third review score
		thirdReviewScore, _ := decimal.NewFromString(rec.Get("averagePoint").String())

		// Third review count
		thirdReviewCount := rec.Get("reviewCount").Int()

		// Space url
		spaceURL := rec.Get("spaceUrl").String()

		// Latitude ..
		latitude, _ := decimal.NewFromString(rec.Get("building.lat").String())

		// Longitude ..
		longitude, _ := decimal.NewFromString(rec.Get("building.lon").String())

		// Get capacity and setup 1 people
		capacity := int16(1)
		if rec.Get("capacity").Int() != 0 {
			capacity = int16(rec.Get("capacity").Int())
		}

		// Get id on platform
		iop := rec.Get("id").Int()

		// Name title + friendlyTitle
		name := rec.Get("friendlyTitle").String() + " " + rec.Get("title").String()

		// Get access
		accesses := rec.Get("building.summaryAccesses.#.access").Array()
		lines := rec.Get("building.summaryAccesses.#.line").Array()
		stations := rec.Get("building.summaryAccesses.#.station").Array()

		size := len(accesses)
		access := ""
		for i := 0; i < size; i++ {
			access = access + lines[i].String() + " " + stations[i].String() + "駅から" + accesses[i].String() + "\n"
		}

		// Get description
		description := rec.Get("seoDescription").String()

		spaces = append(spaces, &briefSpace{
			PlatformId:       2,
			ID:               iop,
			Name:             name,
			Capacity:         capacity,
			ThirdReviewScore: thirdReviewScore,
			ThirdReviewCount: int32(thirdReviewCount),
			HourlyMinPrice:   hourlyMinPrice,
			HourlyMaxPrice:   hourlyMaxPrice,
			SpaceURL:         spaceURL,
			Latitude:         latitude,
			Longitude:        longitude,
			Thumbnails:       thumbnailsArray,
			Access:           access,
			Description:      description,
		})
	}

	return spaces, nil
}

func (s *service) saveSpaces(ctx context.Context, tx *sqlx.Tx, spaces []*briefSpace, oneType api.InstabaseEventType) error {

	if len(spaces) >= 1 {
		for _, space := range spaces {
			res, err := tx.ExecContext(ctx,
				`insert into spaces (
						iop,
						name,
						capacity,
						hourly_min_price,
						hourly_max_price,
						third_review_score,
						third_review_count,
						platform_id,
						space_url,
						latitude,
						longitude,
						thumbnails,
						access,
						created,
						updated,
						description,
						equipment_description,
						amenities,
						event_types
						) values (
						?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?
						) on duplicate key update 
						id = last_insert_id(id),
						updated = values(updated)`,
				space.ID,
				space.Name,
				space.Capacity,
				space.HourlyMinPrice,
				space.HourlyMaxPrice,
				space.ThirdReviewScore,
				space.ThirdReviewCount,
				space.PlatformId,
				space.SpaceURL,
				space.Latitude,
				space.Longitude,
				space.Thumbnails,
				space.Access,
				time.Now(),
				time.Now(),
				space.Description,
				"",
				"",
				"")

			if err != nil {
				s.Log.Error(err)
				return err
			}

			// Update eventTyps
			spaceID, err := res.LastInsertId()
			s.Log.Info("Space  ID:       " + strconv.FormatInt(spaceID, 10))
			s.Log.Info("Space Name:       " + space.Name)
			if err != nil {
				s.Log.Error(err)
				return err
			}

			// Update elastic ..
			var targets []*api.ElasticTarget
			targets = s.Elastic.AddTargets(targets, api.TypeSpace, []int64{spaceID})
			err = s.Elastic.BulkUpdate(ctx, tx, targets)

			if err != nil {
				s.Log.Error(err)
				return err
			}
			_, err = tx.ExecContext(ctx,
				`update instabase_event_types set start_page = ? , hourly_at = now(), updated=now() where id =?`,
				oneType.StartPage+1, oneType.ID)
			if err != nil {
				s.Log.Info(err)
				return err
			}
		}
	} else {
		_, err := tx.ExecContext(ctx, `update instabase_event_types set start_page = ?, hourly_at = now(), updated = now() where id = ?`,
			1, oneType.ID)

		if err != nil {
			s.Log.Info(err)
			return err
		}
	}

	return nil
}
