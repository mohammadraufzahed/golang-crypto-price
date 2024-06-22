package scheduler

import "github.com/jasonlvhit/gocron"

var Scheduler gocron.Scheduler

func Initialize() {
	Scheduler = *gocron.NewScheduler()
}

func Start() {
	go func() {
		<-Scheduler.Start()
	}()
}
