package scraper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/yun313350095/Noonde/api"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Service
type Service struct {
	Conf  api.ConfService
	Log   api.LogService
	MySQL api.MySQLService
}

// NewSpaceMarketClient ..
func (s *Service) NewSpaceMarketClient() (api.SpaceMarketClient, error) {
	pcli, err := s.newProxyClient()
	if err != nil {
		s.Log.Error(err)
		return nil, err
	}

	return &spaceMarketClient{
		service:     s,
		proxyClient: pcli,
	}, nil
}

// NewInstabaseClient ..
func (s *Service) NewInstabaseClient() (api.InstabaseClient, error) {
	pcli, err := s.newProxyClient()
	if err != nil {
		s.Log.Error(err)
		return nil, err
	}

	return &instabaseClient{
		service:     s,
		proxyClient: pcli,
	}, nil
}

func (s *Service) newProxyClient() (*proxyClient, error) {
	var err error

	version := uuid.Must(uuid.NewV4()).String()[0:8]

	auth := s.Conf.String("luminati.username") +
		"-country-jp-session-" +
		version +
		":" +
		s.Conf.String("luminati.password")

	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	proxyStr := s.Conf.String("luminati.url")
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		s.Log.Error(err)
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	transport.ProxyConnectHeader = http.Header{}
	transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
	client := &http.Client{
		Transport: transport,
	}

	return &proxyClient{
		client: client,
	}, nil
}

// --------------------------------------------------------------------------------------------

func (s *Service) buildBody(tmpl string, data interface{}) string {
	var body bytes.Buffer

	template.Must(template.New("tmpl").Parse(tmpl)).Execute(&body, data)

	return body.String()
}

func (s *Service) gql(path string) (string, error) {
	bgql, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	bgql2, err := json.Marshal(string(bgql))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query := string(bgql2)
	query = strings.TrimPrefix(query, `"`)
	query = strings.TrimSuffix(query, `"`)

	return query, nil
}
