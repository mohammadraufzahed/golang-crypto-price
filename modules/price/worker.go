package price

import "fmt"

type PriceJob struct {
	Name string
}

func (job *PriceJob) Execute() error {
	fmt.Println("Got the price for " + job.Name)
	return nil
}
