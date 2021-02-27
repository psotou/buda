package main

import (
	"fmt"
	"os"
	"src/buda-app/buda"
	"strconv"
)

func main() {
	var key string = os.Getenv("KEY")
	var secret string = os.Getenv("SECRET")
	var BTC string = "btc-clp"
	var ETH string = "eth-clp"
	var BCH string = "bch-clp"
	var LTC string = "ltc-clp"

	budaClient, err := buda.NewAPIClient(key, secret)

	if err != nil {
		panic(err)
	}

	marketBTC, _ := budaClient.GetTickerByMarket(BTC)
	marketETH, _ := budaClient.GetTickerByMarket(ETH)
	marketBCH, _ := budaClient.GetTickerByMarket(BCH)
	marketLTC, _ := budaClient.GetTickerByMarket(LTC)

	volumeBTC, _ := budaClient.GetVolumeByMarket(BTC)
	volumeETH, _ := budaClient.GetVolumeByMarket(ETH)
	volumeBCH, _ := budaClient.GetVolumeByMarket(BCH)
	volumeLTC, _ := budaClient.GetVolumeByMarket(LTC)

	// BTC
	priceVar24hBTC, _ := strconv.ParseFloat(marketBTC.PriceVariation24H, 64)
	priceVar7dBTC, _ := strconv.ParseFloat(marketBTC.PriceVariation7D, 64)
	minAskBTC, _ := strconv.ParseFloat(marketBTC.MinAsk[0], 64)
	maxBidBTC, _ := strconv.ParseFloat(marketBTC.MaxBid[0], 64)
	volVentaBTC, _ := strconv.ParseFloat(volumeBTC.AskVolumen24h[0], 64)
	volCompraBTC, _ := strconv.ParseFloat(volumeBTC.BidVolumen24h[0], 64)
	lastPriceBTC, _ := strconv.ParseFloat(marketBTC.LastPrice[0], 64)
	minVentaBTC, _ := strconv.ParseFloat(marketBTC.MinAsk[0], 64)
	maxCompraBTC, _ := strconv.ParseFloat(marketBTC.MaxBid[0], 64)

	// ETH
	priceVar24hETH, _ := strconv.ParseFloat(marketETH.PriceVariation24H, 64)
	priceVar7dETH, _ := strconv.ParseFloat(marketETH.PriceVariation7D, 64)
	minAskETH, _ := strconv.ParseFloat(marketETH.MinAsk[0], 64)
	maxBidETH, _ := strconv.ParseFloat(marketETH.MaxBid[0], 64)
	volVentaETH, _ := strconv.ParseFloat(volumeETH.AskVolumen24h[0], 64)
	volCompraETH, _ := strconv.ParseFloat(volumeETH.BidVolumen24h[0], 64)
	lastPriceETH, _ := strconv.ParseFloat(marketETH.LastPrice[0], 64)
	minVentaETH, _ := strconv.ParseFloat(marketETH.MinAsk[0], 64)
	maxCompraETH, _ := strconv.ParseFloat(marketETH.MaxBid[0], 64)

	// BCH
	priceVar24hBCH, _ := strconv.ParseFloat(marketBCH.PriceVariation24H, 64)
	priceVar7dBCH, _ := strconv.ParseFloat(marketBCH.PriceVariation7D, 64)
	minAskBCH, _ := strconv.ParseFloat(marketBCH.MinAsk[0], 64)
	maxBidBCH, _ := strconv.ParseFloat(marketBCH.MaxBid[0], 64)
	volVentaBCH, _ := strconv.ParseFloat(volumeBCH.AskVolumen24h[0], 64)
	volCompraBCH, _ := strconv.ParseFloat(volumeBCH.BidVolumen24h[0], 64)
	lastPriceBCH, _ := strconv.ParseFloat(marketBCH.LastPrice[0], 64)
	minVentaBCH, _ := strconv.ParseFloat(marketBCH.MinAsk[0], 64)
	maxCompraBCH, _ := strconv.ParseFloat(marketBCH.MaxBid[0], 64)

	// LTC
	priceVar24hLTC, _ := strconv.ParseFloat(marketLTC.PriceVariation24H, 64)
	priceVar7dLTC, _ := strconv.ParseFloat(marketLTC.PriceVariation7D, 64)
	minAskLTC, _ := strconv.ParseFloat(marketLTC.MinAsk[0], 64)
	maxBidLTC, _ := strconv.ParseFloat(marketLTC.MaxBid[0], 64)
	volVentaLTC, _ := strconv.ParseFloat(volumeLTC.AskVolumen24h[0], 64)
	volCompraLTC, _ := strconv.ParseFloat(volumeLTC.BidVolumen24h[0], 64)
	lastPriceLTC, _ := strconv.ParseFloat(marketLTC.LastPrice[0], 64)
	minVentaLTC, _ := strconv.ParseFloat(marketLTC.MinAsk[0], 64)
	maxCompraLTC, _ := strconv.ParseFloat(marketLTC.MaxBid[0], 64)

	// BTC
	var volSumaBTC float64 = volVentaBTC + volCompraBTC
	var volVentaPerBTC float64 = (volVentaBTC / volSumaBTC) * 100
	var volCompraPerBTC float64 = (volCompraBTC / volSumaBTC) * 100
	var spreadBTC float64 = minAskBTC - maxBidBTC
	var priceVar24hPerBTC float64 = priceVar24hBTC * 100
	var priceVar7dPerBTC float64 = priceVar7dBTC * 100

	// ETH
	var volSumaETH float64 = volVentaETH + volCompraETH
	var volVentaPerETH float64 = (volVentaETH / volSumaETH) * 100
	var volCompraPerETH float64 = (volCompraETH / volSumaETH) * 100
	var spreadETH float64 = minAskETH - maxBidETH
	var priceVar24hPerETH float64 = priceVar24hETH * 100
	var priceVar7dPerETH float64 = priceVar7dETH * 100

	// BCH
	var volSumaBCH float64 = volVentaBCH + volCompraBCH
	var volVentaPerBCH float64 = (volVentaBCH / volSumaBCH) * 100
	var volCompraPerBCH float64 = (volCompraBCH / volSumaBCH) * 100
	var spreadBCH float64 = minAskBCH - maxBidBCH
	var priceVar24hPerBCH float64 = priceVar24hBCH * 100
	var priceVar7dPerBCH float64 = priceVar7dBCH * 100

	// LTC
	var volSumaLTC float64 = volVentaLTC + volCompraLTC
	var volVentaPerLTC float64 = (volVentaLTC / volSumaLTC) * 100
	var volCompraPerLTC float64 = (volCompraLTC / volSumaLTC) * 100
	var spreadLTC float64 = minAskLTC - maxBidLTC
	var priceVar24hPerLTC float64 = priceVar24hLTC * 100
	var priceVar7dPerLTC float64 = priceVar7dLTC * 100

	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("                                       MERCADO BUDA                                       ")
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("                         BTC                ETH                BCH               LTC")
	fmt.Printf("Precio última orden    %10.f CLP    %10.f CLP    %10.f CLP    %10.f CLP\n", lastPriceBTC, lastPriceETH, lastPriceBCH, lastPriceLTC)
	fmt.Printf("Min precio de venta    %10.f CLP    %10.f CLP    %10.f CLP    %10.f CLP\n", minVentaBTC, minVentaETH, minVentaBCH, minVentaLTC)
	fmt.Printf("Max precio de compra   %10.f CLP    %10.f CLP    %10.f CLP    %10.f CLP\n", maxCompraBTC, maxCompraETH, maxCompraBCH, maxCompraLTC)
	fmt.Printf("SPREAD                 %10.f CLP    %10.f CLP    %10.f CLP    %10.f CLP\n", spreadBTC, spreadETH, spreadBCH, spreadLTC)
	fmt.Printf("Variación 24h          %10.2f %%      %10.2f %%      %10.2f %%      %10.2f %%\n", priceVar24hPerBTC, priceVar24hPerETH, priceVar24hPerBCH, priceVar24hPerLTC)
	fmt.Printf("Variación 7d           %10.1f %%      %10.1f %%      %10.1f %%      %10.1f %%\n", priceVar7dPerBTC, priceVar7dPerETH, priceVar7dPerBCH, priceVar7dPerLTC)
	fmt.Printf("Volumen venta 24h      %4.1f (%2.1f %%)    %4.1f (%2.1f %%)     %4.1f (%2.1f %%)    %4.1f (%2.1f %%)\n", volVentaBTC, volVentaPerBTC, volVentaETH, volVentaPerETH, volVentaBCH, volVentaPerBCH, volVentaLTC, volVentaPerLTC)
	fmt.Printf("Volumen compra 24h     %4.1f (%2.1f %%)    %4.1f (%2.1f %%)     %4.1f (%2.1f %%)    %4.1f (%2.1f %%)\n", volCompraBTC, volCompraPerBTC, volCompraETH, volCompraPerETH, volCompraBCH, volCompraPerBCH, volCompraLTC, volCompraPerLTC)
	fmt.Println("-------------------------------------------------------------------------------------------")
}
