package price

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/router"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/scheduler"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/worker"
)

var service = PriceService{}

func Initialize() {
	scheduler.Scheduler.Every(10).Seconds().Do(func() {
		prices := service.GetPrices()
		for _, price := range prices {
			fmt.Printf("Symbol: %s, Price: %v\n", price.Symbol, price.Price)
		}
	})
	router.Server.Get("price/job", func(c *fiber.Ctx) error {
		job := &PriceJob{
			Name: "BTCUSDT",
		}
		worker.JobPool.Add(job)
		return c.SendString("Hello")
	})
}
