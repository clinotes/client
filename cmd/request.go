package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type request interface {
	host() string
	post(data interface{}) error
	postScan(data interface{}, result interface{}) error
	url() string
}

type apiResponse struct {
	Error bool `json:"error"`
}

type apiResponseData struct {
	Data interface{}
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

func (r apiRequest) postScan(data interface{}, result interface{}) error {
	buff, err := json.Marshal(data)

	if err != nil {
		return errors.New("Unable to marshal note")
	}

	req, err := http.NewRequest("POST", r.url(), bytes.NewBuffer(buff))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return errors.New("Failed to send request")
	}

	if resp.StatusCode != 200 {
		return errors.New("Received HTTP error code")
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&apiResponseData{result})
	defer resp.Body.Close()

	if err != nil {
		return errors.New("Failed to parse response")
	}

	return nil
}

func (r apiRequest) post(data interface{}) error {
	return r.postScan(data, apiResponse{})
}

func (r apiRequest) host() string {
	if APIHostname != "" {
		return APIHostname
	}

	return FallbackHostname
}
