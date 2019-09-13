package scraper

import (
	"github.com/yun313350095/Noonde/api/scraper/instabase/types"
	"strconv"
	"time"
)

type instabaseClient struct {
	service     *Service
	proxyClient *proxyClient
	retried     int
}

// RefreshProxyClient ..
func (c *instabaseClient) RefreshProxyClient() error {
	s := c.service

	pcli, err := s.newProxyClient()
	if err != nil {
		return err
	}

	c.proxyClient = pcli

	return nil
}

// SearchSpaceList  ..
func (c *instabaseClient) SearchSpaceList(page int, eventType int) (string, error) {
	s := c.service
	var res string
	var err error

	s.Log.Info(eventType)

	gql, err := s.gql(s.Conf.String("goroot") + "/scraper/instabase/gql/search_space_list.gql")
	if err != nil {
		s.Log.Info(err)
		return "", err
	}

	params := &types.SearchRoomsQueryVariables{}
	params.Page = page
	params.UsageIds = eventType
	params.Query = gql

	body := s.buildBody(`{
         "variables": {
           "areaId": null,
           "bottomRightLat": null,
           "bottomRightLon": null,
           "capacityIds": null,
           "categoryIds": null,
           "conditionIds": null,
           "equipmentIds": null,
           "fromDateDay": null,
           "fromDateMonth": null,
           "fromDateYear": null,
           "fromTime": null,
           "orderBy": null,
           "page": {{.Page}},
           "perPage": null,
           "prefectureId": 13,
           "stationId": null,
           "toTime": null,
           "topLeftLat": null,
           "topLeftLon": null,
           "usageIds": [{{.UsageIds}}],
           "wardId": null
           },
          "query": "{{.Query}}"
         }`,
		params)

	res, err = c.proxyClient.postJSON("https://www.instabase.jp/graphql", body, map[string]string{
		"Content-Type": "application/json",
	})

	s.Log.Info(res)

	if err != nil {
		s.Log.Error(err)
		return "", err
	}

	time.Sleep(1 * time.Second)

	return res, nil
}

// Get space info
func (c *instabaseClient) GetSpaceDetails(iop int64) (string, error) {
	s := c.service
	var res string
	var err error

	s.Log.Info("Space iop:  " + strconv.FormatInt(iop, 10))

	gql, err := s.gql(s.Conf.String("goroot") + "/scraper/instabase/gql/space_details.gql")
	if err != nil {
		s.Log.Error(err)
		return "", err
	}

	params := &types.GetSpaceDetails{}
	params.ID = iop
	params.Query = gql

	body := s.buildBody(`{
		"variables": {
           "id": "{{.ID}}"
          },
        "query": "{{.Query}}"
		}`,
		params)

	res, err = c.proxyClient.postJSON("https://www.instabase.jp/graphql", body, map[string]string{
		"Content-Type": "application/json",
	})

	if err != nil {
		s.Log.Info(err)
		return "", err
	}

	time.Sleep(1 * time.Second)

	return res, nil
}
