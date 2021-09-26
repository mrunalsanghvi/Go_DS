package weather

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type failoverWeatherReporter struct {
	url string
}
type primaryWeatherReporter struct {
	url string
}
type WeatherGetter interface {
	GetWeatherClient() (*weatherReportItems, error)
}
type weatherReportItems struct {
	Temp   float64
	Wspeed float64
}

func (pwr *primaryWeatherReporter) GetWeatherClient() (*weatherReportItems, error) {

	urlString := pwr.url
	u, err := url.Parse(urlString)
	if err != nil {

		log.Fatal(err)
		return nil, err
	}
	res, err1 := http.Get(u.String())
	if err1 != nil {
		log.Fatal(err)
		return nil, err1
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err2)
		return nil, err2
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, err2
	}
	jmaps := make(map[string]interface{})
	json.Unmarshal(body, &jmaps)
	reportItems := new(weatherReportItems)
	curval := jmaps["current"].(map[string]interface{})

	for key, val := range curval {

		switch key {
		case "wind_speed":
			reportItems.Wspeed = val.(float64)
		case "temperature":
			reportItems.Temp = val.(float64)
		}
	}
	return reportItems, nil
}

func (fwr *failoverWeatherReporter) GetWeatherClient() (*weatherReportItems, error) {

	urlString := fwr.url
	reportItems := new(weatherReportItems)
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	res, err1 := http.Get(u.String())
	if err1 != nil {
		log.Fatal(err1)
		return nil, err1
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		log.Fatal(err2)
		return nil, err2
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, err2
	}
	jmaps := make(map[string]interface{})
	err3 := json.Unmarshal(body, &jmaps)
	if err != nil {
		log.Fatal(err3)
		return nil, err3
	}
	mainval := jmaps["main"].(map[string]interface{})

	for key, val := range mainval {
		switch key {
		case "temp":
			reportItems.Temp = val.(float64)
		}
	}
	windval := jmaps["wind"].(map[string]interface{})
	for key, val := range windval {
		switch key {
		case "speed":
			reportItems.Wspeed = val.(float64)
		}
	}
	return reportItems, nil
}

func NewWeatherReporters(reporterType string) WeatherGetter {

	switch reporterType {
	case "primary":
		return &primaryWeatherReporter{
			url: "http://api.weatherstack.com/current?access_key=d3dc7541da0dcc0b8677163f6ce9fccb&query=Melbourne",
		}
	case "failover":
		return &failoverWeatherReporter{
			url: "http://api.openweathermap.org/data/2.5/weather?q=Melbourne,Au&appid=252f93844b6d0c0b4d0f04ccf315974b",
		}
	}
	return nil
}

func (w *weatherReportItems) get(h http.ResponseWriter, req *http.Request) {

	reportHandler := NewWeatherReporters("primary")
	wReport, err := reportHandler.GetWeatherClient()
	if err != nil {
		reportHandler := NewWeatherReporters("failover")
		wReport, err = reportHandler.GetWeatherClient()
		if err != nil {
			log.Fatal("Nothing is Okay")
			h.WriteHeader(http.StatusBadRequest)
			h.Write([]byte("Response Not Found"))
		}
	}
	responseJson, err := json.Marshal(wReport)
	if err != nil {
		log.Fatal(err)
		h.WriteHeader(http.StatusBadRequest)
		h.Write([]byte("Response Not Found"))
	}
	h.Header().Add("content-type", "application/json")
	h.WriteHeader(http.StatusOK)
	h.Write(responseJson)
}

func (w *weatherReportItems) Reporters(h http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.get(h, req)
		return
	default:
		h.WriteHeader(http.StatusMethodNotAllowed)
		h.Write([]byte("method not allowed"))
		return
	}
}

func NewWeatherHandlers() *weatherReportItems {
	return &weatherReportItems{}
}
