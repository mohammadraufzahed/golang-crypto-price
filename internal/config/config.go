package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type influxDB struct {
	Url    string
	Token  string
	Org    string
	Bucket string
}

var instance *Config

type Config struct {
	InfluxDB influxDB
}

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load the .env")
	}

	instance = &Config{
		InfluxDB: influxDB{
			Url:    os.Getenv("INFLUXDB_URL"),
			Token:  os.Getenv("INFLUXDB_TOKEN"),
			Org:    os.Getenv("INFLUXDB_ORG"),
			Bucket: os.Getenv("INFLUXDB_BUCKET"),
		},
	}
}

func Get() *Config {
	if instance == nil {
		panic("Config is not initialized")
	}
	return instance
}
