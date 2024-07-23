package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Location, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsRes := Location{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return Location{}, err
		}
		return locationsRes, nil
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationsRes := Location{}
	err = json.Unmarshal(dat, &locationsRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return locationsRes, nil

}

/* func getRequest(id int) string {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + fmt.Sprintf("%d/", id))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatal(err)
	}
	return location.Name

} */
