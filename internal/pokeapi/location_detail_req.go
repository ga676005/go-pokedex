package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetail(idOrName string) (LocationDetailResponse, error) {
	endPoint := "/location-area/" + idOrName
	fullURL := baseURL + endPoint

	rawData, ok := c.cahce.Get(fullURL)
	if ok {
		locationDetailResponse := LocationDetailResponse{}
		json.Unmarshal(rawData, &locationDetailResponse)
		return locationDetailResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationDetailResponse{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetailResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return LocationDetailResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	rawData, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationDetailResponse{}, err
	}

	locationDetailResponse := LocationDetailResponse{}
	err = json.Unmarshal(rawData, &locationDetailResponse)
	if err != nil {
		return LocationDetailResponse{}, err
	}

	c.cahce.Add(fullURL, rawData)

	return locationDetailResponse, nil
}
