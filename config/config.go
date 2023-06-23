package config

import (
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
)

type Config struct {
	InfluxDB InfluxDB
	Slack    Slack
}

type InfluxDB struct {
	Url    string
	Token  string
	Org    string
	Bucket string
}

type Slack struct {
	Token   string
	Channel string
}

var DB influxdb2.Client
var Env Config

func GetEnvironmentVariable() error {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("url_agent.json")
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	Env = config

	return nil
}
