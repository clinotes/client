/**
 * clinot.es client
 * Copyright (C) 2016 Sebastian Müller
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

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
