package utils

import (
	"encoding/json"
	"fmt"
)

type APIResponse struct {
	Datatable Datatable `json:"datatable"`
	Meta      Meta      `json:"meta"`
}

type Meta struct {
	NextCursorId int `json:"next_cursor_id"`
}

type Datatable struct {
	Data    []Data   `json:"data"`
	Columns []Column `json:"columns"`
}

type Data struct {
	Ticker     string  `json:"ticker"`
	Date       string  `json:"date"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	Volume     float64 `json:"volume"`
	Dividend   float64 `json:"dividend"`
	Split      float64 `json:"split"`
	Adj_open   float64 `json:"adj_open"`
	Adj_high   float64 `json:"adj_high"`
	Adj_low    float64 `json:"adj_low"`
	Adj_close  float64 `json:"adj_close"`
	Adj_volume float64 `json:"adj_volume"`
}

// UnmarshalJSON handles unmarshaling the heterogeneous JSON array data provided by the NASDAQ API.
func (n *Data) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.Ticker, &n.Date, &n.Open, &n.High, &n.Low, &n.Close, &n.Volume, &n.Dividend, &n.Split, &n.Adj_open, &n.Adj_high, &n.Adj_low, &n.Adj_close, &n.Volume}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Data: %d != %d", g, e)
	}
	return nil
}

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
