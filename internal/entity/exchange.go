package entity

import (
	"encoding/json"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

const (
	URL = "https://www.cbr-xml-daily.ru/daily_json.js"
)

//Exchange

type Exchange struct {
	Currency string
	Response string
}

type Valute struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"СharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}

type DailyValues struct {
	Data    map[string]interface{}
	Valutes map[string]Valute
}

func (daily *DailyValues) SetValutes() {
	resp, err := http.Get(URL)
	if err != nil {
		//TODO logger
		panic("Implement logger")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&daily.Data)

	if err != nil {
		//TODO logger
		panic("Implement me")
	}

	valute := daily.Data["Valute"].(map[string]interface{})

	daily.Valutes = make(map[string]Valute)

	for key, value := range valute {
		var v Valute
		mapstructure.Decode(value, &v)
		daily.Valutes[key] = v
	}

}
