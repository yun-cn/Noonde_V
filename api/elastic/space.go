package elastic

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic"
	"github.com/shopspring/decimal"
	"github.com/yun313350095/Noonde/api"
	"strconv"
	"strings"
)

func (s *Service) updateSpaces(ctx context.Context, tx *sqlx.Tx, req *elastic.BulkService, spaceIDs []int64) (*elastic.BulkService, error) {
	if len(spaceIDs) == 0 {
		return req, nil
	}

	type space struct {
		ID               int64           `db:"id"`
		Name             string          `db:"name"`
		States           int             `db:"states"`
		City             string          `db:"city"`
		StateText        string          `db:"state_text"`
		Capacity         int             `db:"capacity"`
		Amenities        string          `db:"amenities"`
		EventTypes       string          `db:"event_types"`
		PlatformID       int             `db:"platform_id"`
		Latitude         decimal.Decimal `db:"latitude"`
		Longitude        decimal.Decimal `db:"longitude"`
		ThirdReviewScore decimal.Decimal `db:"third_review_score"`
	}
	var spaces []*space

	// Get spaces.
	query, aa, err := sqlx.In(
		`select id, name, city, state_text, capacity, amenities, event_types, latitude, longitude, third_review_score, states, platform_id from spaces where id in (?)`,
		spaceIDs,
	)
	if err != nil {
		s.Log.Error(err)
		return req, nil
	}
	err = tx.SelectContext(ctx, &spaces, query, aa...)
	if err != nil {
		s.Log.Error(err)
		return req, nil
	}

	spaceFor := map[int64]*space{}
	for _, space := range spaces {
		spaceFor[space.ID] = space
	}

	for _, spaceID := range spaceIDs {
		if space, ok := spaceFor[spaceID]; ok {

			// Update.
			var ss, amenitisArray, eventTypesArray, tt []string

			if len(space.StateText) != 0 {
				ss = append(ss, space.StateText)
			}

			if len(space.City) != 0 {
				ss = append(ss, space.City)
			}

			ss = append(ss, space.Name)

			// Amenities to array
			if len(space.Amenities) != 0 {
				amenities := strings.TrimPrefix(space.Amenities, `[`)
				amenities = strings.TrimSuffix(amenities, `]`)
				amenities = strings.Replace(amenities, `"`, "", -1)
				amenitisArray = strings.Split(amenities, `,`)
			}

			// Event type to array
			if len(space.EventTypes) != 0 {
				eventTypes := strings.TrimPrefix(space.EventTypes, `[`)
				eventTypes = strings.TrimSuffix(eventTypes, `]`)
				eventTypes = strings.Replace(eventTypes, `"`, "", -1)
				eventTypesArray = strings.Split(eventTypes, `,`)
			}

			// Set space platform id
			tt = append(tt, "kind-"+strconv.Itoa(space.PlatformID))

			// Set space state
			if space.States == 1 {
				tt = append(tt, "states-1")
			}

			if space.States == 2 {
				tt = append(tt, "states-2")
			}

			doc := &struct {
				Capacity   int             `json:"capacity"`
				Longitude  decimal.Decimal `json:"longitude"`
				Latitude   decimal.Decimal `json:"latitude"`
				Tags       []string        `json:"tags"`
				Amenities  []string        `json:"amenities"`
				EventTypes []string        `json:"event_types"`
				Review     decimal.Decimal `json:"review"`
				Type       string          `json:"type"`
				Search     string          `json:"search"`
				ID         int64           `json:"id"`
			}{
				space.Capacity,
				space.Longitude,
				space.Latitude,
				tt,
				amenitisArray,
				eventTypesArray,
				space.ThirdReviewScore,
				api.TypeSpace,
				strings.Join(ss, " | "),
				space.ID,
			}

			q := elastic.NewBulkIndexRequest().
				Index(s.Conf.String("elastic.index")).
				Type(s.Conf.String("elastic.type")).
				Id(api.TypeSpace + "-" + strconv.FormatInt(space.ID, 10)).
				Doc(doc)
			req = req.Add(q)
			continue

		} else {

			// Delete.
			q := elastic.NewBulkDeleteRequest().
				Index(s.Conf.String("elastic.index")).
				Type(s.Conf.String("elastic.type")).
				Id(api.TypeSpace + "-" + strconv.FormatInt(spaceID, 10))
			req = req.Add(q)
		}

	}

	return req, nil
}
