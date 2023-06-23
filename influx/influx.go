package influx

import (
	"context"
	"log"
	"mayf_url_monitor/config"
	"mayf_url_monitor/model"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InfluxInit() {
	config.DB = influxdb2.NewClient(config.Env.InfluxDB.Url, config.Env.InfluxDB.Token)
}

func URLStatToInflux(response model.Response) {
	if response.Name != "" {
		writeAPI := config.DB.WriteAPIBlocking(config.Env.InfluxDB.Org, config.Env.InfluxDB.Bucket)
		p := influxdb2.NewPointWithMeasurement(response.Name).
			AddTag("url", response.Url).
			AddField("status_code", response.StatusCode).
			AddField("response_time", response.ResponseTime).
			SetTime(time.Now())
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Println("Cant Write DATA in InfluxDB with ", err, response)
		} else {
			log.Println("Write DATA in InfluxDB", response)
		}
	} else {
		log.Println("Cant Write DATA in InfluxDB with No Name DATA", response)
	}
}
