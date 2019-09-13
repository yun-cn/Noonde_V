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

func (s *service) getSpaceDetails(client api.InstabaseClient, iop int64) (*spaceDetails, error) {
	var err error
	var body string
	var data gjson.Result

	// Get space details
	body, err = client.GetSpaceDetails(iop)
	if err != nil {
		s.Log.Info(err)
		panic(err)
	}

	// Response body to hash
	hashByte := sha256.Sum256([]byte(body))
	hash := hex.EncodeToString(hashByte[:])

	s.Log.Info(hash)

	// Edited  hash at
	hashAt := time.Now()

	// Parse space details
	data = gjson.Parse(body)

	if len(data.Get("data.rooms").Array()) == 0 {
		space := &spaceDetails{
			PlatformId: 2,
			States:     2,
			IOP:        iop,
			Hash:       hash,
			HashAt:     &hashAt,
		}

		return space, nil
	}

	record := data.Get("data.rooms").Array()[0]

	// Hourly min price
	hourlyMinPriceStr := record.Get("summaryMinPrice").String()
	hourlyMinPrice, _ := decimal.NewFromString(hourlyMinPriceStr)
	s.Log.Info(hourlyMinPrice)

	// Hourly max price
	hourlyMaxPriceStr := record.Get("summaryMaxPrice").String()
	hourlyMaxPrice, _ := decimal.NewFromString(hourlyMaxPriceStr)
	s.Log.Info(hourlyMaxPrice)

	// Third review core
	thirdReviewScore, _ := decimal.NewFromString(record.Get("averagePoint").String())
	s.Log.Info(thirdReviewScore)

	// Third review count
	thirdReviewCount := record.Get("reviewCount").Int()
	s.Log.Info(thirdReviewCount)

	//  Space url
	spaceURL := record.Get("spaceUrl").String()
	s.Log.Info(spaceURL)

	// Space uid
	uid := record.Get("uid").String()
	s.Log.Info(uid)

	// Space name
	name := record.Get("friendlyTitle").String() + " " + record.Get("title").String()
	s.Log.Info(name)

	// Space capacity and setup 1 people
	capacity := int16(1)
	if record.Get("capacity").Int() != 0 {
		capacity = int16(record.Get("capacity").Int())
	}

	// Space latitude
	latitude, _ := decimal.NewFromString(record.Get("building.lat").String())
	s.Log.Info(latitude)

	// Space longitude
	longitude, _ := decimal.NewFromString(record.Get("building.lon").String())
	s.Log.Info(longitude)

	// Thumbnails array
	thumbnails := record.Get("images.#.largeFilePath")
	thumbnailList := make([]string, len(thumbnails.Array()))

	for i, thumbnail := range thumbnails.Array() {
		thumbnailList[i] = thumbnail.String()
	}

	thumbnailsArray, _ := json.Marshal(thumbnailList)

	// Amenities array
	amenitiesBase := map[string]string{
		"鏡":            "full_length_mirror",
		"Wi-Fi(無線LAN)": "wifi",
		"椅子":           "chairs",
		"楽器":           "performance",
		"有線LAN":        "internet_hikari",
		"ピアノ":          "piano", // 没有
		"駐車場":          "parking",
		"テレビ":          "tv",
		"パソコン":         "computer", // 没有
		"モニター":         "monitor",
		"キッチン":         "kitchen_facilities",
		"テーブル":         "tables",
		"エアコン":         "aircon",
		"DVDプレイヤー":     "dvd_br_player",
		"その他の備品":       "Others",  // 没有
		"マイクセット":       "mic_set", // 没有
		"ホワイトボード":      "whiteboard",
		"プロジェクター":      "projector",
		"有線マイクセット":     "wired_mic_set", // 没有
		"アンプ・スピーカー":    "speaker",       // 没有
		"文具類一式":        "stationery",    // 没有
	}

	amenities := record.Get("freeEquipments.#.title")
	amenityList := []string{}

	if len(amenities.Array()) != 0 {
		for _, v := range amenities.Array() {
			if val, ok := amenitiesBase[v.String()]; ok {
				//amenityList[i] = val
				amenityList = append(amenityList, val)
			} else {
				continue
			}
		}
	}
	s.Log.Info(amenityList)
	amenityArray, _ := json.Marshal(amenityList)

	// Get access
	accesses := record.Get("building.summaryAccesses.#.access").Array()
	lines := record.Get("building.summaryAccesses.#.line").Array()
	stations := record.Get("building.summaryAccesses.#.station").Array()

	size := len(accesses)
	access := ""
	for i := 0; i < size; i++ {
		access = access + lines[i].String() + " " + stations[i].String() + "駅から" + accesses[i].String() + "\n"
	}

	// Get description
	description := record.Get("seoDescription").String()

	// Get equipment description
	equipmentDescription := record.Get("notice").String()

	// Get id on platform
	idOnPlatform := record.Get("id").Int()

	// Get space address
	address := record.Get("building.address").String()

	// Get space state_text
	addressArray := strings.Split(address, " ")
	stateText := addressArray[1]

	// Get space city
	city := addressArray[2]

	// Get space event_types
	eventTypesBase := map[string]string{
		"勉強会":      "class",
		"セラピー":     "other",
		"作業場所":     "office",
		"パーティー":    "party",
		"面接・試験":    "office",
		"撮影・収録":    "photo_shoot",
		"ヨガ・ダンス":   "sports",
		"女子会・ママ会":  "party",
		"ボードゲーム":   "party",
		"ワークショップ":  "social_event",
		"カウンセリング":  "social_event",
		"上映会・映画鑑賞": "film_shoot",
		"セミナー・研修":  "class",
		"打ち合わせ・商談": "class",
		"レッスン・講座":  "class",
	}

	eventTypes := record.Get("usages.#.title")
	eventTypeList := []string{}

	if len(eventTypes.Array()) != 0 {
		for _, v := range eventTypes.Array() {
			if val, ok := eventTypesBase[v.String()]; ok {
				eventTypeList = append(eventTypeList, val)
			} else {
				continue
			}
		}
	}

	eventTypeList = removeDuplicates(eventTypeList)

	eventTypesArray, _ := json.Marshal(eventTypeList)

	space := &spaceDetails{
		PlatformId:           2,
		States:               1,
		Name:                 name,
		UID:                  uid,
		IOP:                  idOnPlatform,
		Capacity:             capacity,
		ThirdReviewScore:     thirdReviewScore,
		ThirdReviewCount:     int32(thirdReviewCount),
		HourlyMinPrice:       hourlyMinPrice,
		HourlyMaxPrice:       hourlyMaxPrice,
		SpaceURL:             spaceURL,
		Latitude:             latitude,
		Longitude:            longitude,
		Thumbnails:           thumbnailsArray,
		Access:               access,
		Description:          description,
		Amenities:            amenityArray,
		EquipmentDescription: equipmentDescription,
		StateText:            stateText,
		EventTypes:           eventTypesArray,
		City:                 city,
		Address:              address,
		Hash:                 hash,
		HashAt:               &hashAt,
	}

	return space, nil
}

func (s *service) saveSpaceDetails(ctx context.Context, tx *sqlx.Tx, space *spaceDetails) error {

	// Space hash dose not exits or not same ?!
	var hash string

	err := tx.GetContext(ctx, &hash, `select hash from spaces where iop = ? and platform_id =2`, space.IOP)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	if (hash != space.Hash && space.States == 1) || (len(hash) == 0 && space.States == 1) {

		res, err := tx.ExecContext(ctx, `update spaces set 
				uip = ?,
			    name = ?,
      			capacity = ?,
		        third_review_score = ?,
      	        third_review_count = ?,
      			hourly_min_price = ?,
		        hourly_max_price = ?,
		        space_url = ?,
		        latitude = ?,
                longitude = ?,
                thumbnails = ?,
                amenities = ?,
                address = ?,
                description = ?,
                state_text =?,
			    city = ?,
                states = ?,
                equipment_description = ?,
                event_types = ?,
                hash = ?,
                hash_at= ?,
                updated = now() where platform_id = ? and iop = ?`,
			space.UID,
			space.Name,
			space.Capacity,
			space.ThirdReviewScore,
			space.ThirdReviewCount,
			space.HourlyMinPrice,
			space.HourlyMaxPrice,
			space.SpaceURL,
			space.Latitude,
			space.Longitude,
			space.Thumbnails,
			space.Amenities,
			space.Address,
			space.Description,
			space.StateText,
			space.City,
			space.States,
			space.EquipmentDescription,
			space.EventTypes,
			space.Hash,
			space.HashAt,
			space.PlatformId,
			space.IOP,
		)

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
		err = tx.GetContext(ctx, &spaceID, `select id from spaces where iop = ? and platform_id = 2`,
			space.IOP)
		s.Log.Info("Space ID:    " + strconv.FormatInt(spaceID, 10))

		if err != nil {
			s.Log.Error(err)
			return err
		}

		// Update elastic_search
		var targets []*api.ElasticTarget
		targets = s.Elastic.AddTargets(targets, api.TypeSpace, []int64{spaceID})
		err = s.Elastic.BulkUpdate(ctx, tx, targets)
		if err != nil {
			s.Log.Error(err)
			return err
		}

	} else {
		res, err := tx.ExecContext(ctx,
			`update spaces set states = ?, hash = ?, hash_at = ? where iop = ? and platform_id = 2`,
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
		err = tx.GetContext(ctx, &spaceID, `select id from spaces where iop = ? and platform_id = 2`,
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

// removeDuplicates ..
func removeDuplicates(elements []string) []string { // change string to int here if required
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{} // change string to int here if required
	result := []string{}             // change string to int here if required

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
