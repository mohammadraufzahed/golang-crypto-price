package price

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/influxdb"
)

var (
	priceService PriceService
)

type TickerPrice struct {
	Symbol string  `json:"symbol"`
	Price  float32 `json:"price,string"`
}

type PriceService struct{}

func (p *PriceService) GetPrices() []TickerPrice {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price")

	if err != nil {
		fmt.Println("Failed to get the response")
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Failed to read the body: %v", err)
		return nil
	}

	var prices []TickerPrice
	err = json.Unmarshal(body, &prices)

	if err != nil {
		log.Fatalf("Failed to parse the body: %v", err)
		return nil
	}

	return prices
}

func (p *PriceService) SyncPrices(prices []TickerPrice) {
	influxdb := influxdb.Get()
	for _, price := range prices {
		tags := map[string]string{"symbol": price.Symbol}
		fields := map[string]interface{}{"price": price.Price}

		p := influxdb2.NewPoint("price", tags, fields, time.Now())
		influxdb.WriteAPI.WritePoint(p)
	}

	influxdb.WriteAPI.Flush()
}

func initServices() {
	priceService = PriceService{}
}
