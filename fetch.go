package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

var ErrMissingCookies = errors.New("twilio/gen: missing value for TWILIO_COOKIES env var")

type Service struct {
	GoName         string
	FriendlyName   string
	IsBeta         bool
	EndpointGroups []EndpointGroup
}

type EndpointGroup struct {
	GoName       string
	FriendlyName string
	Location     string
}

type Endpoint struct {
	GoName       string
	FriendlyName string
	Location     string
	Method       string
	DocLink      string
}

func FetchAPIInfo(api string) (map[string]interface{}, error) {
	body, err := fetch(api)
	return body.(map[string]interface{}), err
}

func FetchEndpointGroupInfo(api string) ([]interface{}, error) {
	body, err := fetch(api)
	return body.([]interface{}), err
}

func fetch(api string) (interface{}, error) {
	cookies := os.Getenv("TWILIO_COOKIES")
	if cookies == "" {
		return nil, ErrMissingCookies
	}
	url := fmt.Sprintf("https://www.twilio.com/console/api/api-explorer%s?relative=true", api)
	//log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.twilio.com/console/runtime/api-explorer/voice/calls")
	req.Header.Set("Cookie", cookies)
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result interface{}
	return result, json.NewDecoder(resp.Body).Decode(&result)
}
