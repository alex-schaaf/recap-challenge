package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const baseUrl = "https://data.nasdaq.com/api/v3/datatables/QUOTEMEDIA/PRICES"

// FetchSymbolData from the the NASDAQ End of Day US Stock Prices API.
func FetchSymbolData(symbol, start, end string) (APIResponse, error) {
	url := getUrl(symbol, start, end)
	response, err := http.Get(url)
	if err != nil {
		return APIResponse{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return APIResponse{}, err
	}

	apiResponse, err := parseResponseBody(body)
	if err != nil {
		return apiResponse, err
	}

	if len(apiResponse.Datatable.Data) == 0 {
		return apiResponse, errors.New("No data")
	}

	return apiResponse, nil
}

func getUrl(symbol, start, end string) string {
	q := url.Values{}
	q.Add("date.gte", start)
	q.Add("date.lte", end)
	q.Add("ticker", symbol)
	q.Add("api_key", os.Getenv("API_KEY"))

	return baseUrl + "?" + q.Encode()
}

func parseResponseBody(body []byte) (APIResponse, error) {
	var dat APIResponse
	if err := json.Unmarshal(body, &dat); err != nil {
		return dat, err
	}
	return dat, nil
}
