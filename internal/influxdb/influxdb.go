package influxdb

import (
	"sync"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/mohammadraufzahed/golang-crypto-price/internal/config"
)

type influxDB struct {
	Client   influxdb2.Client
	WriteAPI api.WriteAPI
	QueryAPI api.QueryAPI
}

var (
	instance *influxDB
	once     sync.Once
)

func initialize() {
	config := config.Get()
	once.Do(func() {
		client := influxdb2.NewClient(config.InfluxDB.Url, config.InfluxDB.Token)
		writeApi := client.WriteAPI(config.InfluxDB.Org, config.InfluxDB.Bucket)
		queryApi := client.QueryAPI(config.InfluxDB.Org)

		instance = &influxDB{
			Client:   client,
			WriteAPI: writeApi,
			QueryAPI: queryApi,
		}
	})
}

func Get() *influxDB {
	if instance == nil {
		initialize()
	}
	return instance
}
