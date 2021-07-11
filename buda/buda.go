package buda

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	BaseURL              = "https://www.buda.com/api/v2"
	MarketTickerEndpoint = "/markets/%s/ticker"
	VolumeEndpoint       = "/markets/%s/volume"
	BalancesEndpoint     = "/balances"
	ElementsPerPage      = "20"
)

type APIClient struct {
	Key    string
	Secret string
	Client *http.Client
}

type Ticker struct {
	MarketID          string   `json:"market_id"`
	LastPrice         []string `json:"last_price"`
	MinAsk            []string `json:"min_ask"`
	MaxBid            []string `json:"max_bid"`
	Volume            []string `json:"volume"`
	PriceVariation24H string   `json:"price_variation_24h"`
	PriceVariation7D  string   `json:"price_variation_7d"`
}

type TickerSingle struct {
	Ticker Ticker `json:"ticker"`
}

type Volume struct {
	AskVolumen24h []string `json:"ask_volume_24h"`
	AskVolumen7d  []string `json:"ask_volume_7d"`
	BidVolumen24h []string `json:"bid_volume_24h"`
	BidVolumen7d  []string `json:"bid_volume_7d"`
	MarketId      string   `json:"market_id"`
}

type VolumeSingle struct {
	Volume Volume `json:"volume"`
}

type Metadata struct {
	CurrentPage int `json:"current_page"`
	TotalCount  int `json:"total_count"`
	TotalPages  int `json:"total_pages"`
}

type Balance struct {
	ID     string   `json:"id"`
	Amount []string `json:"amount"`
	// AvailableAmount       []string `json:"available_amount"`
	// FrozenAmount          []string `json:"frozen_amount"`
	// PendingWithdrawAmount []string `json:"pending_withdraw_amount"`
}

type Balances struct {
	Balances []Balance `json:"balances"`
}

type BalanceSingle struct {
	Balance Balance `json:"balance"`
}

func (client *APIClient) SignRequest(params ...string) string {
	h := hmac.New(sha512.New384, []byte(client.Secret))
	h.Write([]byte(strings.Join(params, " ")))
	return hex.EncodeToString(h.Sum(nil))
}

func (client *APIClient) AuthenticatedRequest(request *http.Request) (*http.Request, error) {
	var signature string
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano()*1e6, 10)

	switch request.Method {
	case "POST":
		{
			var body []byte
			body, err := io.ReadAll(request.Body)
			if err != nil {
				return nil, err
			}
			signature = client.SignRequest(request.Method, request.URL.RequestURI(), base64.StdEncoding.EncodeToString(body), timestamp)
		}
	case "GET":
		{
			signature = client.SignRequest(request.Method, request.URL.RequestURI(), timestamp)
		}
	}

	request.Header.Set("X-SBTC-APIKEY", client.Key)
	request.Header.Set("X-SBTC-NONCE", timestamp)
	request.Header.Set("X-SBTC-SIGNATURE", signature)

	return request, nil
}

func NewAPIClient(apiKey string, apiSecret string) (*APIClient, error) {
	return &APIClient{Client: &http.Client{}, Key: apiKey, Secret: apiSecret}, nil
}

func (client *APIClient) FormatResource(resource string) string {
	return fmt.Sprintf("%s%s", BaseURL, resource)
}

func (client *APIClient) Get(resource string, private bool) ([]byte, error) {
	req, err := http.NewRequest("GET", client.FormatResource(resource), nil)
	// req, err := http.Get(client.FormatResource(resource))
	if err != nil {
		return nil, err
	}

	if private {
		req, err = client.AuthenticatedRequest(req)
		if err != nil {
			return nil, err
		}
	}

	response, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (client *APIClient) GetTickerByMarket(marketId string) (*Ticker, error) {
	var ticker TickerSingle

	data, err := client.Get(fmt.Sprintf(MarketTickerEndpoint, marketId), false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &ticker)
	if err != nil {
		return nil, err
	}

	return &ticker.Ticker, nil
}

func (client *APIClient) GetVolumeByMarket(marketId string) (*Volume, error) {
	var volume VolumeSingle

	data, err := client.Get(fmt.Sprintf(VolumeEndpoint, marketId), false)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &volume)
	if err != nil {
		return nil, err
	}

	return &volume.Volume, nil
}

func (client *APIClient) GetBalances() ([]Balance, error) {
	var balances Balances

	data, err := client.Get(BalancesEndpoint, true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &balances)
	if err != nil {
		return nil, err
	}

	return balances.Balances, nil
}

func (client *APIClient) GetBalanceByCurrency(currency string) (*Balance, error) {
	var balance BalanceSingle

	data, err := client.Get(BalancesEndpoint+"/"+currency, true)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &balance)
	if err != nil {
		return nil, err
	}

	return &balance.Balance, nil
}
