package main

import (
	"fmt"
	"net/http"

	"github.com/mrunalsanghvi/Go_DS/pkg/weather"
)

func main() {
	weatherHandler := weather.NewWeatherHandlers()
	fmt.Println("starting the http server")
	http.HandleFunc("/v1/weather", weatherHandler.Reporters)
	http.ListenAndServe(":8080", nil)
}
