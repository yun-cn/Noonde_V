package scraper

import (
	"github.com/yun313350095/Noonde/api/scraper/spacemarket/types"
	"time"
)

type spaceMarketClient struct {
	service     *Service
	proxyClient *proxyClient
	retried     int
}

// RefreshProxyClient ..
func (c *spaceMarketClient) RefreshProxyClient() error {
	s := c.service

	pcli, err := s.newProxyClient()
	if err != nil {
		return err
	}

	c.proxyClient = pcli

	return nil
}

// SearchRoomList ..
func (c *spaceMarketClient) SearchRoomListHourly(page int, perPage int, location string, eventType string) (string, error) {
	s := c.service
	var res string
	var err error
	s.Log.Info("Event Type: " + eventType + "\n")

	gql, err := s.gql(s.Conf.String("goroot") + "/scraper/spacemarket/gql/search_room_list.gql")
	if err != nil {
		s.Log.Error(err)
		return "", err
	}
	apiKey := s.Conf.String("spacemarket.x_api_key")

	params := &types.SearchRoomsQueryVariables{}
	params.Page = page
	params.PerPage = perPage
	params.EventType = eventType
	params.PriceType = "HOURLY"
	params.RentType = "DAY_TIME"
	params.Query = gql
	params.Location = "-"

	body := s.buildBody(`{
                "operationName":"Search",
                "variables":{
                  "page":{{.Page}},
                  "perPage":{{.PerPage}},
                  "eventType":"{{.EventType}}",
                  "priceType":"{{.PriceType}}",
                  "location":"{{.Location}}",
                  "rentType":"{{.RentType}}",
                  "maxCapacity":null,
                  "minCapacity": null,
                  "maxPrice":null,
                  "minPrice":null
                },
               "query": "{{.Query}}"}`,
		params,
	)

	res, err = c.proxyClient.postJSON("https://v3api.spacemarket.com/graphql", body, map[string]string{
		"x-api-key": apiKey,
	})

	if err != nil {
		s.Log.Error(err)
		return "", err
	}
	time.Sleep(1 * time.Second)

	return res, nil
}

// Space info
func (c *spaceMarketClient) GetRoomDetails(uid string, rentType string) (string, error) {
	s := c.service
	var res string
	var err error

	s.Log.Info("Room uid: " + uid)

	gql, err := s.gql(s.Conf.String("goroot") + "/scraper/spacemarket/gql/room_details.gql")
	if err != nil {
		s.Log.Error(err)
		return "", err
	}

	params := &types.RoomInfo{}
	params.UID = uid
	params.RentType = rentType
	params.Query = gql

	apiKey := s.Conf.String("spacemarket.x_api_key")

	body := s.buildBody(`{
		 "operationName":"room",
		 "variables":{
           "uid":"{{.UID}}",
           "rentType":"{{.RentType}}",
           "canRentType":"STAY"
		 },
         "query": "{{.Query}}"}`,
		params,
	)

	res, err = c.proxyClient.postJSON("https://v3api.spacemarket.com/graphql", body, map[string]string{
		"x-api-key": apiKey,
	})

	if err != nil {
		s.Log.Error(err)
		return "", err
	}

	return res, nil
}
