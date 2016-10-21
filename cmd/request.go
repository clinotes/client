package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type request interface {
	host() string
	post(data interface{}) (*APIErrorResponse, error)
	url() string
}

type apiRequest struct {
	Action string
}

func newRequest(action string) request {
	return apiRequest{action}
}

func (r apiRequest) url() string {
	return r.host() + r.Action
}

func (r apiRequest) post(data interface{}) (*APIErrorResponse, error) {
	buff, err := json.Marshal(data)

	if err != nil {
		fail("Unable to marshal note")
	}

	req, err := http.NewRequest("POST", r.url(), bytes.NewBuffer(buff))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Failed to send request")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Received HTTP error code")
	}

	var res APIErrorResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	defer resp.Body.Close()

	if err != nil {
		return nil, errors.New("Failed to parse response")
	}

	return &res, nil
}

func (r apiRequest) host() string {
	if APIHostname != "" {
		return APIHostname
	}

	return FallbackHostname
}
