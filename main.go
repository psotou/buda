package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Ticker struct with the amrrket info
type Ticker struct {
	MarketID	string	 `json:"market_id"`
	LastPrice	[]string `json:"last_price"`
	MinAsk	[]string `json:"min_ask"`
	MaxBid	[]string `json:"max_bid"`
	Volume	[]string `json:"volume"`
	PriceVariation24H	string   `json:"price_variation_24h"`
	PriceVariation7D	string   `json:"price_variation_7d"`	
}

// TickerSingle creates a Ticker of type Ticker (which contains what's is int he above struct)
type TickerSingle struct {
	Ticker Ticker `json:"ticker"`
}

func main() {

	const url = "https://www.buda.com/api/v2/markets/btc-clp/ticker.json"

	budaClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := budaClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	
	ticker := TickerSingle{}
	jsonErr := json.Unmarshal(body, &ticker)
	if err != nil {
		fmt.Println(jsonErr)
		return
	}

	var oneDayVariationStr string = ticker.Ticker.PriceVariation24H
	var sevenDayVariationStr string = ticker.Ticker.PriceVariation7D

	oneDayVariation, _ := strconv.ParseFloat(oneDayVariationStr, 64)
	sevenDayVariation, _ := strconv.ParseFloat(sevenDayVariationStr, 64)

	var oneDayVariationPercent float64 = oneDayVariation * 100
	var sevenDayVariationPercent float64 = sevenDayVariation * 100

	fmt.Printf("Precio última transacción	$ %s\n", ticker.Ticker.LastPrice[0])
	fmt.Printf("Menor precio de venta	$ %s\n", ticker.Ticker.MinAsk[0])
	fmt.Printf("Máximo precio de compra	$ %s\n", ticker.Ticker.MaxBid[0])
	fmt.Printf("Variación últimas 24h	%2.2f %%\n", oneDayVariationPercent)
	fmt.Printf("Variación últimos 7d	%2.2f %%\n", sevenDayVariationPercent)
	fmt.Printf("Volumen criptomenda		%s\n", ticker.Ticker.Volume[0])
}