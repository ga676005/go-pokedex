package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	endPoint := "/location-area?offset=0&limit=20"
	fullURL := baseURL + endPoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	rawData, ok := c.cahce.Get(fullURL)
	if ok {
		locationAreaResponse := LocationAreaResponse{}
		json.Unmarshal(rawData, &locationAreaResponse)
		return locationAreaResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	rawData, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(rawData, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cahce.Add(fullURL, rawData)

	return locationAreaResponse, nil
}

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
