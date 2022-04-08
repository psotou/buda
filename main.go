package main

import (
	"buda/buda"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	key    = os.Getenv("BUDA_KEY")
	secret = os.Getenv("BUDA_SECRET")
	BTC    = "btc-clp"
	ETH    = "eth-clp"
	BCH    = "bch-clp"
	LTC    = "ltc-clp"
)

func main() {
	budaClient, err := buda.NewAPIClient(key, secret)

	if err != nil {
		log.Fatal(err.Error())
	}

	marketBTC, _ := budaClient.GetTickerByMarket(BTC)
	marketETH, _ := budaClient.GetTickerByMarket(ETH)
	marketBCH, _ := budaClient.GetTickerByMarket(BCH)
	marketLTC, _ := budaClient.GetTickerByMarket(LTC)

	// BTC
	priceVar24hBTC, _ := strconv.ParseFloat(marketBTC.PriceVariation24H, 64)
	priceVar7dBTC, _ := strconv.ParseFloat(marketBTC.PriceVariation7D, 64)
	lastPriceBTC, _ := strconv.ParseFloat(marketBTC.LastPrice[0], 64)
	minVentaBTC, _ := strconv.ParseFloat(marketBTC.MinAsk[0], 64)
	maxCompraBTC, _ := strconv.ParseFloat(marketBTC.MaxBid[0], 64)

	// ETH
	priceVar24hETH, _ := strconv.ParseFloat(marketETH.PriceVariation24H, 64)
	priceVar7dETH, _ := strconv.ParseFloat(marketETH.PriceVariation7D, 64)
	lastPriceETH, _ := strconv.ParseFloat(marketETH.LastPrice[0], 64)
	minVentaETH, _ := strconv.ParseFloat(marketETH.MinAsk[0], 64)
	maxCompraETH, _ := strconv.ParseFloat(marketETH.MaxBid[0], 64)

	// BCH
	priceVar24hBCH, _ := strconv.ParseFloat(marketBCH.PriceVariation24H, 64)
	priceVar7dBCH, _ := strconv.ParseFloat(marketBCH.PriceVariation7D, 64)
	lastPriceBCH, _ := strconv.ParseFloat(marketBCH.LastPrice[0], 64)
	minVentaBCH, _ := strconv.ParseFloat(marketBCH.MinAsk[0], 64)
	maxCompraBCH, _ := strconv.ParseFloat(marketBCH.MaxBid[0], 64)

	// LTC
	priceVar24hLTC, _ := strconv.ParseFloat(marketLTC.PriceVariation24H, 64)
	priceVar7dLTC, _ := strconv.ParseFloat(marketLTC.PriceVariation7D, 64)
	lastPriceLTC, _ := strconv.ParseFloat(marketLTC.LastPrice[0], 64)
	minVentaLTC, _ := strconv.ParseFloat(marketLTC.MinAsk[0], 64)
	maxCompraLTC, _ := strconv.ParseFloat(marketLTC.MaxBid[0], 64)

	// BTC
	priceVar24hPerBTC := priceVar24hBTC * 100
	priceVar7dPerBTC := priceVar7dBTC * 100

	// ETH
	priceVar24hPerETH := priceVar24hETH * 100
	priceVar7dPerETH := priceVar7dETH * 100

	// BCH
	priceVar24hPerBCH := priceVar24hBCH * 100
	priceVar7dPerBCH := priceVar7dBCH * 100

	// LTC
	priceVar24hPerLTC := priceVar24hLTC * 100
	priceVar7dPerLTC := priceVar7dLTC * 100

	p := message.NewPrinter(language.English)
	p.Printf("%25s %11s %9s %8s\n", "BTC", "ETH", "BCH", "LTC")
	p.Printf("Última orden   %6.f   %6.f   %6.f   %6.f \n", lastPriceBTC, lastPriceETH, lastPriceBCH, lastPriceLTC)
	p.Printf("Min venta      %6.f   %6.f   %6.f   %6.f \n", minVentaBTC, minVentaETH, minVentaBCH, minVentaLTC)
	p.Printf("Max compra     %6.f   %6.f   %6.f   %6.f \n", maxCompraBTC, maxCompraETH, maxCompraBCH, maxCompraLTC)
	p.Printf("Variación 24h  %8.2f %%  %8.2f %%  %6.2f %%  %5.2f %%\n", priceVar24hPerBTC, priceVar24hPerETH, priceVar24hPerBCH, priceVar24hPerLTC)
	p.Printf("Variación 7d   %8.1f %%  %8.1f %%  %6.1f %%  %5.1f %%\n", priceVar7dPerBTC, priceVar7dPerETH, priceVar7dPerBCH, priceVar7dPerLTC)
}
