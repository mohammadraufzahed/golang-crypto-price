package price

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
