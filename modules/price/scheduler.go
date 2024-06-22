package price

import (
	"github.com/mohammadraufzahed/golang-crypto-price/internal/scheduler"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/worker"
)

func initScheduler() {
	scheduler.Scheduler.Every(5).Seconds().Do(func() {
		prices := priceService.GetPrices()
		worker.JobPool.Add(&SyncPrice{
			prices: prices,
		})
	})
}
