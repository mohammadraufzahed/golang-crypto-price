package price

type SyncPrice struct {
	prices []TickerPrice
}

func (job *SyncPrice) Execute() error {
	priceService.SyncPrices(job.prices)
	return nil
}
