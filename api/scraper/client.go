package scraper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type proxyClient struct {
	client *http.Client
}

func (p *proxyClient) get(urlStr string, headers map[string]string) (string, error) {
	parsed, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	request, err := http.NewRequest("GET", parsed.String(), nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for key, val := range headers {
		request.Header.Add(key, val)
	}
	fmt.Println("-> Get " + urlStr)
	response, err := p.client.Do(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

func (p *proxyClient) postJSON(urlStr string, body string, headers map[string]string) (string, error) {
	parsed, err := url.Parse(urlStr)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	request, err := http.NewRequest("POST", parsed.String(), bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	for key, val := range headers {
		request.Header.Add(key, val)
	}
	fmt.Println("-> POST: --> " + urlStr)

	response, err := p.client.Do(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(data), nil
}
