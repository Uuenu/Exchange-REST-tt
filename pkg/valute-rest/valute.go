package valuterest

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	URL = "https://www.cbr-xml-daily.ru/daily_json.js"
)

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

var Valutes DailyValues

func (daily *DailyValues) PickRandomValute() Valute {
	var mu sync.RWMutex
	mu.RLock()
	defer mu.RUnlock()
	for _, v := range daily.Valutes {
		return v
	}
	return Valute{}
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

func RunValuteRest() {
	for {
		Valutes.SetValutes()
		time.Sleep(time.Hour * 12)
	}

}
