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
    volVenta, _ := strconv.ParseFloat(volume.AskVolumen24h[0], 64)
    volCompra, _ := strconv.ParseFloat(volume.BidVolumen24h[0], 64)

    var volSuma float64 = volVenta + volCompra
    var porcentajeVolVenta float64 = (volVenta / volSuma) * 100
    var porcentajeVolCompra float64 = (volCompra / volSuma) * 100
	var spread float64 = minAsk - maxBid
	var oneDayVariationPercent float64 = oneDayVariation * 100
	var sevenDayVariationPercent float64 = sevenDayVariation * 100
    var criptomoneda string = strings.ToUpper(os.Args[1])

    fmt.Println("---------------------------------------------")
    fmt.Printf("                MERCADO %s-CLP\n", criptomoneda)
    fmt.Println("---------------------------------------------")

	fmt.Printf("Precio última orden ejecutada  $ %s\n", market.LastPrice[0])
	fmt.Printf("Menor precio de venta          $ %s\n", market.MinAsk[0])
	fmt.Printf("Máximo precio de compra        $ %s\n", market.MaxBid[0])
	fmt.Printf("SPREAD                         $ %.1f\n", spread)
	fmt.Printf("Variación últimas 24h          %2.2f %%\n", oneDayVariationPercent)
	fmt.Printf("Variación últimos 7d           %2.2f %%\n", sevenDayVariationPercent)
    fmt.Printf("Volumen venta últimas 24h      %3.2f (%2.1f %%)\n", volVenta, porcentajeVolVenta)
    fmt.Printf("Volumen compra últimas 24h     %3.2f (%2.1f %%)\n", volCompra, porcentajeVolCompra)
}
