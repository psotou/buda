package main

import (
	"fmt"
	"os"
	"src/buda-app/buda"
	"strconv"
    "strings"
)

func main() {
	var key string = os.Getenv("KEY")
	var secret string = os.Getenv("SECRET")
	var crypto string = os.Args[1] + "-clp"

	budaClient, err := buda.NewAPIClient(key, secret)

	if err != nil {
		panic(err)
	}

	market, _ := budaClient.GetTickerByMarket(crypto)
	volume, _ := budaClient.GetVolumeByMarket(crypto)

	oneDayVariation, _ := strconv.ParseFloat(market.PriceVariation24H, 64)
	sevenDayVariation, _ := strconv.ParseFloat(market.PriceVariation7D, 64)
	minAsk, _ := strconv.ParseFloat(market.MinAsk[0], 64)
	maxBid, _ := strconv.ParseFloat(market.MaxBid[0], 64)

	var spread float64 = minAsk - maxBid
	var oneDayVariationPercent float64 = oneDayVariation * 100
	var sevenDayVariationPercent float64 = sevenDayVariation * 100

    fmt.Println("---------------------------------------------")
    fmt.Printf("                MERCADO %s-CLP\n", strings.ToUpper(os.Args[1]))
    fmt.Println("---------------------------------------------")

	fmt.Printf("Precio última orden ejecutada  $ %s\n", market.LastPrice[0])
	fmt.Printf("Menor precio de venta          $ %s\n", market.MinAsk[0])
	fmt.Printf("Máximo precio de compra        $ %s\n", market.MaxBid[0])
	fmt.Printf("SPREAD                         $ %.1f\n", spread)
	fmt.Printf("Variación últimas 24h          %2.2f %%\n", oneDayVariationPercent)
	fmt.Printf("Variación últimos 7d           %2.2f %%\n", sevenDayVariationPercent)
	// fmt.Printf("Volumen criptomenda            %s\n", market.Volume[0])
	fmt.Printf("Volumen venta últimas 24h      %s\n", volume.AskVolumen24h[0])
	fmt.Printf("Volumen compra últimas 24h     %s\n", volume.BidVolumen24h[0])
	// fmt.Printf("Volumen venta últimos 7d       %s\n", volume.AskVolumen7d[0])
	// fmt.Printf("Volumen compra últimos 7d      %s\n", volume.BidVolumen7d[0])
}
