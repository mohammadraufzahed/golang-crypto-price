package main

import (
	"github.com/mohammadraufzahed/golang-crypto-price/internal/config"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/influxdb"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/router"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/scheduler"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/worker"
	"github.com/mohammadraufzahed/golang-crypto-price/modules"
)

func main() {
	// Load and initialize the internals
	config.Load()
	worker.InitWorkerPool()
	router.Initialize()
	scheduler.Initialize()
	influxdb.Initialize()
	defer influxdb.Close()

	// Register the modules
	modules.Initialize()

	// Starts
	scheduler.Start()
	router.Start()
}
