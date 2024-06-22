package price_test

import (
	"testing"

	"github.com/mohammadraufzahed/golang-crypto-price/modules/price"
)

func BenchmarkFetchTickerPrices(b *testing.B) {
	service := price.PriceService{}
	for i := 0; i < b.N; i++ {
		service.GetPrices()
	}
}
