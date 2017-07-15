package main

import (
	"fmt"
	"github.com/SimpleApplicationsOrg/stock/alphavantage"
	"log"
)

func main() {

	avClient, err := alphavantage.NewAVClient()
	if err != nil {
		log.Fatalf("main: %s", err.Error())
	}

	data, err := avClient.TimeSeriesIntraday("GOOGL", "1min")
	if err != nil {
		log.Printf("main: %s", err.Error())
		return
	}
	metaData := *data.MetaData
	fmt.Println(metaData.Information(), "->", metaData.OutputSize())
	fmt.Println(metaData.LastRefreshed(), metaData.TimeZone())
	fmt.Println("Time Series", metaData.Interval())

	timeSeries := data.TimeSeries
	for _, key := range timeSeries.Keys() {
		value := (*timeSeries)[key]
		fmt.Println(key, "->", value.Open(), value.High(), value.Low(), value.Close(), value.Volume())
	}
}
