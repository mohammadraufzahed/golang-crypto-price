package main

import (
	"github.com/mohammadraufzahed/golang-crypto-price/internal/router"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/scheduler"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/worker"
	"github.com/mohammadraufzahed/golang-crypto-price/modules"
)

func main() {
	worker.InitWorkerPool()
	router.Initialize()
	scheduler.Initialize()
	modules.Initialize()
	scheduler.Start()
	router.Start()
}
